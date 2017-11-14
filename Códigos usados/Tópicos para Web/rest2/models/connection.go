package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getConnection() *sql.DB {
	dsn := "root:dieguim961100@/gocrud?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
