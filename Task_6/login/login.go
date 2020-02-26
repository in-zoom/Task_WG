package login

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connect() *sql.DB {
    databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}