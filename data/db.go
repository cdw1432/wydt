package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBInit() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		fmt.Println("DB: connection failed")
	} else {
		DB = db
		fmt.Println("DB: connected")
	}
}
