package models

import (
	"database/sql"
	"log"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetAllUsers(db *sql.DB) []User {
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func GetUserById(db *sql.DB, id string) User {
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(users) == 1 {
		return users[0]
	}
	return User{}
}
func CreateUser(db *sql.DB, username string, password string) int {

	createdAt := time.Now()

	result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(id)
}
