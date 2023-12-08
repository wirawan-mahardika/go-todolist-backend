package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:wm050604@tcp(localhost:3306)/todolist?parseTime=true")
	if err != nil { panic(err) }

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxIdleTime(time.Hour)
	
	return db
}