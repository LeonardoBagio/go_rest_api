package database

import (
	"database/sql"

	"example.com/rest-api/database/table"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	table.CreateTableUser(DB)
	table.CreateTableEvent(DB)
}
