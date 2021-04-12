package headline

import (
	"crawler/db"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

type HeadlineNewsCrawler struct {
	DB *sqlx.DB
	Client *http.Client
}

func (c HeadlineNewsCrawler) Crawling() error {
	// Get Naver headline news list
	log.Print("Try Get Naver headline news list...")
	urls, err := scrapUrls(c.Client)
	if err != nil {
		fmt.Errorf("Failed to scrapUrls : %v", err)
		return err
	}
	log.Println("Success")

	// result logging & check already exist
	hash := sha256.New()
	for _, url := range urls {
		log.Println("=> " + url)

		hash.Reset()
		hash.Write([]byte(url))
		hashedValue := hash.Sum(nil)
		hashedValueString := hex.EncodeToString(hashedValue)

		nid, err := db.SelectNewsByHash(c.DB, hashedValueString)
		if err == sql.ErrNoRows {
			// new article!
			article, err := scrapNews(c.Client, url)
			if err != nil {
				return err
			}

			newArticle := db.News{
				//Id:         0,
				Title:      article.Title,
				Body:       article.Body,
				Hash:       hashedValueString,
				CreateTime: article.CreateTime,
				UpdateTime: article.UpdateTime,
			}
			//spew.Dump(newArticle)
			//log.Println(newArticle.CreateTime.UTC())
			nid, err = db.InsertNewArticle(c.DB, newArticle)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err

		}

		// check last crawled comment
		lc, err := db.GetLastCrawledComment(c.DB, nid)
		if err == sql.ErrNoRows {
			lc.CreateTime = time.Unix(0, 0)
		} else if err != nil {
			return err
		}

		// scrap new comment
		scrapedComments, err := scrapComments(c.Client, url, lc.CreateTime)
		if err != nil {
			return err
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

			spew.Dump(newComment)
			if _, err := db.InsertNewComment(c.DB, newComment); err != nil {
				return err
			}
		}
	}

	return nil
}