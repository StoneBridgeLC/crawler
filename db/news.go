package db

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type News struct {
	Id int64 `db:"id"`
	Title string `db:"title"`
	Body string `db:"body"`
	Hash string `db:"hash"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func SelectNewsByHash(db *sqlx.DB, hashValue string) (int64, error) {
	var id int64
	err := db.QueryRowx(`
select n.id
from news n 
where n.hash = ?;`, hashValue).Scan(&id)
	return id, err
}

func InsertNewArticle(db *sqlx.DB, news News) (int64, error) {
	result, err := db.NamedExec(`
		insert into news (title, body, hash, create_time, update_time)
		values (:title, :body, :hash, :create_time, :update_time);`, news)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}