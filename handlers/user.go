package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"maxposch/simplegobackend/models"

	"github.com/gorilla/mux"
)

// Example curl -X POST http://localhost:80/user/ -d "username=johndoe" -d "password=secret"
func UserPostHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")
		id := models.CreateUser(db, username, password)

		json.NewEncoder(w).Encode(id)
	}
}

func UserGetAllHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := models.GetAllUsers(db)

		json.NewEncoder(w).Encode(users)
	}
}

func UserGetByIDHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		user := models.GetUserById(db, id)

		json.NewEncoder(w).Encode(user)
	}
}
