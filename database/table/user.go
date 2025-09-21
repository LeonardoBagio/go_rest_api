package table

import "database/sql"

func CreateTableUser(DB *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create users table")
	}
}
