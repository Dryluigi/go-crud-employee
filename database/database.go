package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dsn := "root@tcp(localhost:3306)/crud-employee"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
