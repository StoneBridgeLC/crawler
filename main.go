package main

import (
	"crawler/service/naver/headline"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	DB_USER := os.Getenv("DB_USER")
	DB_PW := os.Getenv("DB_PW")
	DB_IP := os.Getenv("DB_IP")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USER, DB_PW, DB_IP, DB_PORT, DB_NAME)
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		sugar.Fatal("failed to db open", err)
	}
	defer db.Close()
	sugar.Infow("DB open",
		"dataSourceName", dataSourceName)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defaultTransportPointer, ok := http.DefaultTransport.(*http.Transport)
	if !ok { panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport")) }
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	client := &http.Client{Transport: &defaultTransport}

	crawler := headline.HeadlineNewsCrawler{
		Log: sugar,
		DB: db,
		Client: client,
	}

	if err := crawler.Crawling(); err != nil {
		crawler.Log.Errorw("failed to Crawling()",
			"error with stack", errors.WithStack(err))
	}
}