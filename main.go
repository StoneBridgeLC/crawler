package main

import (
	"crawler/service/naver/headline"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	crawler := headline.HeadlineNewsCrawler{}

	db, err := sqlx.Open("mysql", "root:qwer1234@tcp(localhost:3306)/post_analyzer?parseTime=true")
	if err != nil {
		log.Fatalf("Failed to db open : %v", err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	crawler.DB = db

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: &defaultTransport}

	crawler.Client = client

	if err := crawler.Crawling(); err != nil {
		log.Println(err)
	}
}