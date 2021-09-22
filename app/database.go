package app

import (
	"database/sql"
	"github.com/chwlr/golang-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_category_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
