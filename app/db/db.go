package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB = initDB()

func initDB() *sqlx.DB {
	db, err := connectDB()
	if err != nil {
		log.Fatal("connect error:", err)
	}
	return db
}

func connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", "root", "password", "api_db"))
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("Ping failure:%s", err.Error())
	}

	return db, nil
}
