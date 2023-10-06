package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(){
	var err error
	connStr := "root:@tcp(localhost:3306)/projectdb"
	DB, err = sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}
}