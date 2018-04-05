package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Context struct {
	DB *sql.DB
}

func DBInitial(*sql.DB) *sql.DB {
	var err error
	db, err := sql.Open("mysql", "toor:toor@/DBSelling")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewContext() *Context {
	return nil
}
