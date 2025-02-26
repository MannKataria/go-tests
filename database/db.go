package database

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func InsertUser(db *sql.DB, name string) (int, error) {
	query := "INSERT INTO users (name) VALUES (?)"
	result, err := db.Exec(query, name)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func GetUser(db *sql.DB, id int) (User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	var user User

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}
	return user, nil
}
