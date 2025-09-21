package table

import "database/sql"

func CreateTableRegistration(DB *sql.DB) {
	createTable := `
	CREATE TABLE IF NOT EXISTS registration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`

	_, err := DB.Exec(createTable)

	if err != nil {
		panic("Could not create registration table")
	}
}
