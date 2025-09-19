package models

import (
	"example.com/rest-api/database"
	"example.com/rest-api/utils"
)

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

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	return err
}
