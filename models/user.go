package models

import "example.com/rest-api/database"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}
