package Mean

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type cats struct {
	meanTail     float32
	meanWhiskers float32
}

func Mean(db *sql.DB) {
	rows, err := db.Query("SELECT SUM(cats.tail_length) / COUNT(cats.name), SUM(cats.whiskers_length) / COUNT(cats.name) FROM cats")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]*cats, 0)
	for rows.Next() {
		bk := new(cats)
		err := rows.Scan(&bk.meanTail, &bk.meanWhiskers)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range bks {
		//fmt.Println(bk.meanTail, bk.meanWhiskers)
		_, err := db.Exec("INSERT INTO cats_stat (tail_length_mean, whiskers_length_mean) VALUES ($1, $2)", bk.meanTail, bk.meanWhiskers)
		if err != nil {
			panic(err)
		}
	}
}
