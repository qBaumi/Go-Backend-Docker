package main

import (
	"database/sql"
	"fmt"
	"log"
	"maxposch/simplegobackend/handlers"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userRouter := r.PathPrefix("/user").Subrouter()

	/* if you have a mariadb server but port 3306 is not allowed through firewall
	you have to tunnel with ssh like this:
	ssh -N -L 3306:127.0.0.1:3306 root@yourserverip
	*/
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	// init
    { 
        query := `
            CREATE TABLE users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

        if _, err := db.Exec(query); err != nil {
            log.Fatal(err)
        }
    }
	userRouter.HandleFunc("", handlers.UserPostHandler(db)).Methods("POST")
	userRouter.HandleFunc("", handlers.UserGetAllHandler(db)).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.UserGetByIDHandler(db)).Methods("GET")
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
