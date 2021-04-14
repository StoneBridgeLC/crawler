package headline

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"moul.io/http2curl"
	"net/http"
	"strings"
	"time"
)

type Article struct {
	Url string
	Title string
	Body string
	CreateTime time.Time
	UpdateTime time.Time
}

func (c HeadlineNewsCrawler) scrapNews(url string) (Article, error) {
	c.Log.Infow("start scrapNews",
		"url", url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Article{}, err
	}
	req.Header.Set("Authority", "news.naver.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		curlCmd, err := http2curl.GetCurlCommand(req)
		if err != nil {
			return Article{}, errors.Wrap(err, "response statuscode is not 200 and build curl command fail.")
		}
		return Article{}, errors.New(fmt.Sprintf("response statuscode is not 200. curl : %s", curlCmd.String()))
	}

	// response content-type=text/html:charset=EUC-KR
	reader := transform.NewReader(resp.Body, korean.EUCKR.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return Article{}, err
	}

	var (
		timeParsingError error = nil
		articleTimes []time.Time = make([]time.Time, 0, 2)
	)
	doc.Find("span.t11").Each(func(i int, selection *goquery.Selection) {
		t, err := parseTime(selection.Text())
		if err != nil {
			timeParsingError = err
			return
		}
		articleTimes = append(articleTimes, t)
	})

	if timeParsingError != nil {
		return Article{}, timeParsingError
	}

	article := Article{
		Url: url,
		CreateTime: articleTimes[0],
		//UpdateTime: articleTimes[0],
	}
	// 입력시간만 있고 수정시간이 없는 경우도 있음
	if len(articleTimes) > 1 {
		article.UpdateTime = articleTimes[1]
	} else {
		article.UpdateTime = articleTimes[0]
	}

	article.Title = doc.Find("h3#articleTitle").Text()
	doc.Find("#articleBodyContents").Children().Remove()
	article.Body = strings.TrimSpace(doc.Find("#articleBodyContents").Text())

	c.Log.Infow("success scrapNews")
	return article, nil
}

// layout : 2006.01.02. 오후 3:04
func parseTime(timestr string) (time.Time, error) {
	kstTimeZone := time.FixedZone("KST", 9*60*60)

	if t, err := time.ParseInLocation("2006.01.02. 오전 3:04", timestr, kstTimeZone); err == nil {
		return t, nil
	}

	if t, err := time.ParseInLocation("2006.01.02. 오후 3:04", timestr, kstTimeZone); err == nil {
		return t.Add(time.Hour * 12), nil
	}

	return time.Time{}, fmt.Errorf("unknown layout \"%s\"", timestr)
}