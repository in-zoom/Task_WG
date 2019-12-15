package main

import (
	"backend/group_count"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var db *sql.DB

func main() {
	e := godotenv.Load(".env") //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	group_count.GroupCount(db)
}
