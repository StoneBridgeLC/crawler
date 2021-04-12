package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Comment struct {
	Id int64 `db:"id"`
	NId int64 `db:"nid"`
	PId sql.NullInt64 `db:"pid"`
	Body string `db:"body"`
	IsPos sql.NullBool `db:"isPos"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func GetLastCrawledComment(db *sqlx.DB, nid int64) (Comment, error) {
	c := Comment{}
	err := db.QueryRowx(
		`select c.id, nid, body, pid, isPos, create_time, update_time
from comments c
where c.nid = ?
order by c.create_time desc 
limit 1;`, nid).StructScan(&c)
	return c, err
}

func InsertNewComment(db *sqlx.DB, c Comment) (int64, error) {
	result, err := db.NamedExec(
		`insert into comments (nid, body, pid, isPos, create_time, update_time)
VALUES (:nid, :body, :pid, :isPos, :create_time, :update_time);`, c)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}