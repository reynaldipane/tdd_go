package main

import (
	"database/sql"
	"fmt"
)

type user struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name", db:"name"`
	Age  int    `json:"age" db:"age"`
}

func (u *user) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func (u *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)
	_, err := db.Query(statement)

	return err
}

func (u *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	_, err := db.Query(statement)

	return err
}

func (u *user) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users (name, age) VALUES ('%s', %d)", u.Name, u.Age)
	_, err := db.Query(statement)

	return err
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	users := []user{}
	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}