package headline

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_getNewsList(t *testing.T) {

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: &defaultTransport}

	urls, err := getNewsList(client)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(urls[i])
	}
}
