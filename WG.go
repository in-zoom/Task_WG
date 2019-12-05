package main

import (
	"database/sql"
	"fmt"
	"os"

	"Backend/Mean"
	"Backend/Median"
	"Backend/ModeTail"
	"Backend/ModeWhiskers"
	"github.com/joho/godotenv"
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
		fmt.Println("ошибка")
		os.Exit(1)
	}

	ModeTail.ModeTail(db)
	ModeWhiskers.ModeWhiskers(db)
	Mean.Mean(db)
	Median.Median(db)

}
