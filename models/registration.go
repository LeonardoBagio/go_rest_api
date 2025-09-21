package models

import (
	"database/sql"

	"example.com/rest-api/database"
)

type Registration struct {
	ID      int64
	EventID int64 `binding:"required"`
	UserID  int64 `binding:"required"`
}

func GetAllRegistration() ([]Registration, error) {
	query := `SELECT * FROM registration`
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	registrations := []Registration{}

	for rows.Next() {
		var registration Registration
		err := rows.Scan(&registration.ID, &registration.EventID, &registration.UserID)

		if err != nil {
			return nil, err
		}

		registrations = append(registrations, registration)
	}

	return registrations, nil
}

func GetRegistrationById(id int64) (*Registration, error) {
	query := `SELECT * FROM registration WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var registration Registration
	err := row.Scan(&registration.ID, &registration.EventID, &registration.UserID)

	if err != nil {
		return nil, err
	}

	return &registration, nil
}

func Register(event *Event, userId int64) (result sql.Result, err error) {
	query := `INSERT INTO registration (event_id, user_id) VALUES (?, ?)`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err = statement.Exec(event.ID, userId)

	return result, err
}

func CancelRegistration(event *Event, userId int64) error {
	query := `DELETE FROM registration WHERE event_id = ? AND user_id = ?`
	statement, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(event.ID, userId)

	return err
}
