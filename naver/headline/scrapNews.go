package headline

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
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

func scrapNews(client *http.Client, url string) (Article, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Article{}, err
	}
	req.Header.Set("Authority", "news.naver.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return Article{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Article{}, fmt.Errorf("scrapNews : response statuscode is not 200\n%v", spew.Sdump(req))
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
		UpdateTime: articleTimes[1],
	}
	article.Title = doc.Find("h3#articleTitle").Text()
	doc.Find("#articleBodyContents").Children().Remove()
	article.Body = strings.TrimSpace(doc.Find("#articleBodyContents").Text())

	return article, nil
}

// layout : 2006.01.02. 오후 3:04
func parseTime(timestr string) (time.Time, error) {
	kstTimeZone := time.FixedZone("KST", int(time.Hour * 9))

	if t, err := time.ParseInLocation("2006.01.02. 오전 3:04", timestr, kstTimeZone); err == nil {
		return t, nil
	}

	if t, err := time.ParseInLocation("2006.01.02. 오후 3:04", timestr, kstTimeZone); err == nil {
		return t.Add(time.Hour * 12), nil
	}

	return time.Time{}, fmt.Errorf("parseTime : unknown layout \"%s\"", timestr)
}