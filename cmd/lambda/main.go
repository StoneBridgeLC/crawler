package main

import (
	"context"
	"crawler/service/naver/headline"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, event events.CloudWatchEvent) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()
	sugar := logger.Sugar().Named(event.ID)

	DB_USER := os.Getenv("DB_USER")
	DB_PW := os.Getenv("DB_PW")
	DB_IP := os.Getenv("DB_IP")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USER, DB_PW, DB_IP, DB_PORT, DB_NAME)
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return errors.Wrap(err, "failed to db open")
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

	crawler := &headline.HeadlineNewsCrawler{
		Log: sugar,
		DB: db,
		Client: client,
	}

	if err := crawler.Crawling(); err != nil {
		crawler.Log.Errorw("failed to Crawling()",
			"error with stack", errors.WithStack(err))
	}

	return nil
}
