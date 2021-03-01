package implement_repository_pattern

import (
	"database/sql"
	"time"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:ichsanasya@tcp(localhost:3306)/belajar_golang_mysql?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(19 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	// defer db.Close()
	return db
}