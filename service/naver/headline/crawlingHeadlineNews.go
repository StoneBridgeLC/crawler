package headline

import (
	"crawler/db"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"time"
	"unicode/utf8"
)

// bodyLenLowestLimit 이하의 길이의 뉴스는 크롤링하지 않는다.
// 기사에 뉴스 동영상만 올려놓고 내용 안쓰는 기사 거르기 위함.
const bodyLenLowestLimit int = 200

type HeadlineNewsCrawler struct {
	Log    *zap.SugaredLogger
	DB     *sqlx.DB
	Client *http.Client
}

func (c HeadlineNewsCrawler) Crawling() error {
	c.Log.Info("start Crawling")
	errorIdx := 0

	// Get Naver headline news list
	urls, err := c.scrapUrls()
	if err != nil {
		return err
	}

	// result logging & check already exist
	hash := sha256.New()
	for index, url := range urls {
		hash.Reset()
		hash.Write([]byte(url))
		hashedValue := hash.Sum(nil)
		hashedValueString := hex.EncodeToString(hashedValue)

		c.Log.Infow("task",
			"task index", index,
			"target url", url,
			"hash", hashedValueString)

		nid, err := db.SelectNewsByHash(c.DB, hashedValueString)
		if err == sql.ErrNoRows {
			// new article!
			c.Log.Info("new article!")
			article, err := c.scrapNews(url)
			if err != nil {
				errorIdx++
				c.Log.Errorw("failed to scrapNews",
					"url", url,
					"errorIdx", errorIdx,
					"error with stack", errors.WithStack(err))
				continue
			}

			// bodyLenLowestLimit 이하의 길이의 뉴스는 크롤링하지 않는다.
			// 기사에 뉴스 동영상만 올려놓고 내용 안쓰는 기사 거르기 위함.
			if utf8.RuneCountInString(article.Body) <= bodyLenLowestLimit {
				c.Log.Infow("skip this article. body size is too short",
					"url", url,
					"body", article.Body)
				continue
			}

			newArticle := db.News{
				//Id:         0,
				Title: article.Title,
				Body:  article.Body,
				Hash:  hashedValueString,
				Url: sql.NullString{
					String: url,
					Valid:  true,
				},
				CreateTime: article.CreateTime,
				UpdateTime: article.UpdateTime,
			}
			nid, err = db.InsertNewArticle(c.DB, newArticle)
			if err != nil {
				errorIdx++
				c.Log.Errorw("failed to insert new scraped news",
					"failed article", newArticle,
					"errorIdx", errorIdx,
					"error with stack", errors.WithStack(err))
				continue
			}
		} else if err != nil {
			errorIdx++
			c.Log.Errorw("failed to select news by hash value",
				"hash value", hashedValueString,
				"errorIdx", errorIdx,
				"error with stack", errors.WithStack(err))
			continue
		}

		// check last crawled comment
		lc, err := db.GetLastCrawledComment(c.DB, nid)
		if err == sql.ErrNoRows {
			lc.CreateTime = time.Unix(0, 0)
		} else if err != nil {
			errorIdx++
			c.Log.Errorw("failed to select last crawled comment",
				"nid", nid,
				"errorIdx", errorIdx,
				"error with stack", errors.WithStack(err))
			continue
		}

		// scrap new comment
		scrapedComments, err := c.scrapComments(url, lc.CreateTime)
		if err != nil {
			errorIdx++
			c.Log.Errorw("failed to scrapComments",
				"url", url,
				"standard time", lc.CreateTime,
				"errorIdx", errorIdx,
				"error with stack", errors.WithStack(err))
			continue
		}

		for _, scrapedComment := range scrapedComments {
			newComment := db.Comment{
				//Id:         0,
				NId:        nid,
				PId:        sql.NullInt64{},
				Body:       scrapedComment.Body,
				IsPos:      sql.NullBool{},
				CreateTime: scrapedComment.CreateTime,
				UpdateTime: scrapedComment.UpdateTime,
			}

			newcid, err := db.InsertNewComment(c.DB, newComment)
			if err != nil {
				errorIdx++
				c.Log.Errorw("failed to insert new scraped comment",
					"failed comment", newComment,
					"errorIdx", errorIdx,
					"error with stack", errors.WithStack(err))
				continue
			}
			c.Log.Infow("insert new comment",
				"comment id", newcid)
		}
	}

	c.Log.Infow("finish Crawling",
		"occurred error count", errorIdx)
	return nil
}
