package main

import (
	"Backend/Mean"
	"Backend/Median"
	"Backend/Mode"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"os"
)

type statistic struct {
	Mean   float32
	Median float32
	Mode   []float32
}

var db *sql.DB
var columTail string = "cats.tail_length"
var columWhiskers string = "cats.whiskers_length"

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
	tailStatistic := readStatistic(db, columTail)
	whiskersStatistic := readStatistic(db, columWhiskers)
	ins := "INSERT INTO cats_stat (tail_length_mean, whiskers_length_mean, tail_length_median, whiskers_length_median, tail_length_mode, whiskers_length_mode) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = db.Exec(ins, tailStatistic.Mean, whiskersStatistic.Mean, tailStatistic.Median, whiskersStatistic.Median, pq.Array(tailStatistic.Mode), pq.Array(whiskersStatistic.Mode))
	//_, err = db.Exec("INSERT INTO cats_stat (tail_length_mean, whiskers_length_mean, tail_length_median, whiskers_length_median, tail_length_mode, whiskers_length_mode) VALUES ($1, $2, $3, $4, $5, $6)", tailStatistic.Mean, whiskersStatistic.Mean, tailStatistic.Median, whiskersStatistic.Median, tailStatistic.Mode, whiskersStatistic.Mode)
	if err != nil {
		panic(err)
	}
}
func readStatistic(db *sql.DB, columName string) *statistic {
	ResultStatistic := new(statistic)
	ResultStatistic.Mean = Mean.Mean(db, columName)
	ResultStatistic.Median = Median.Median(db, columName)
	ResultStatistic.Mode = Mode.Mode(db, columName)
	return ResultStatistic
}
