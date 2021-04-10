package headline

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const naverNewsDomain = "https://news.naver.com"

var defaultClient *http.Client
func init() {
	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	defaultClient = &http.Client{Transport: &defaultTransport}
}


// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// curl 'https://news.naver.com/' \
//   -H 'authority: news.naver.com' \
//   -H 'sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"' \
//   -H 'sec-ch-ua-mobile: ?0' \
//   -H 'upgrade-insecure-requests: 1' \
//   -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36' \
//   -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' \
//   -H 'sec-fetch-site: none' \
//   -H 'sec-fetch-mode: navigate' \
//   -H 'sec-fetch-user: ?1' \
//   -H 'sec-fetch-dest: document' \
//   -H 'accept-language: ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7' \
//   -H 'cookie: NNB=RTJJKMVHRBYGA; paneOrderNewsHome=today_main_news%2Csection_politics%2Csection_economy%2Csection_society%2Csection_life%2Csection_world%2Csection_it' \
//   --compressed

func getNewsList(client *http.Client) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, naverNewsDomain, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "news.naver.com")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"89\", \"Chromium\";v=\"89\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Accept-Language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cookie", "NNB=RTJJKMVHRBYGA; paneOrderNewsHome=today_main_news%2Csection_politics%2Csection_economy%2Csection_society%2Csection_life%2Csection_world%2Csection_it")


	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	newsUrls := make([]string, 0)
	doc.Find(".hdline_article_tit a").Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if ok {
			newsUrls = append(newsUrls, url)
		}
	})

	return newsUrls, nil
}
