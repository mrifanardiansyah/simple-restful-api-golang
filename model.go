package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(name, age) values('%s', %d)", u.Name, u.Age)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM users where id=%d", u.ID)

	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
}

func (u *User) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users set name='%s', age=%d where id=%d", u.Name, u.Age, u.ID)

	_, err := db.Exec(statement)

	return err
}

func (u *User) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("delete from users where id=%d", u.ID)
	_, err := db.Exec(statement)

	return err
}

func getUsers(db *sql.DB, start, count int) ([]User, error) {
	statement := fmt.Sprintf("SELECT id, name, age from users where limit %d offset %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
