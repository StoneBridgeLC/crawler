package headline

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"testing"
)

func Test_scrapComments(t *testing.T) {
	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: &defaultTransport}

	comments, err := scrapComments(client, "https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=104&oid=015&aid=0004528764", SetLastCrawledTimeisNull())
	if err != nil {
		t.Error(err)
	}

	spew.Dump(comments)
}
