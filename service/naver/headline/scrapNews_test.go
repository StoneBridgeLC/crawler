package headline

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_scrapNews(t *testing.T) {
	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: &defaultTransport}

	article, err := scrapNews(client, "https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=104&oid=015&aid=0004528764")
	if err != nil {
		t.Error(err)
	}

	spew.Dump(article)
}

func Test_parseTime(t *testing.T) {
	type args struct {
		timestr string
	}
	tests := []struct {
		name string
		args args
		want time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "오후 테스트1",
			args: args{timestr: "2021.04.12. 오후 2:48"},
			want: time.Date(2021, time.April, 12, 14, 48, 0, 0, time.FixedZone("KST", int(time.Hour * 9))),
			wantErr: false,
		},
		{
			name: "오후 테스트2",
			args: args{timestr: "2021.04.12. 오후 11:48"},
			want: time.Date(2021, time.April, 12, 23, 48, 0, 0, time.FixedZone("KST", int(time.Hour * 9))),
			wantErr: false,
		},
		{
			name: "오전 테스트1",
			args: args{timestr: "2021.04.12. 오전 2:48"},
			want: time.Date(2021, time.April, 12, 2, 48, 0, 0, time.FixedZone("KST", int(time.Hour * 9))),
			wantErr: false,
		},
		{
			name: "에러 테스트",
			args: args{timestr: "2021.04.12 오전 2:48"},
			want: time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTime(tt.args.timestr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}