package table

import "database/sql"

func CreateTableEvent(DB *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create events table")
	}
}
