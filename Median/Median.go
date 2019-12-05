package Median

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type cats struct {
	medianTail     float32
	medianWhiskers float32
}

func Median(db *sql.DB) {
	rows, err := db.Query("SELECT percentile_disc(0.5) WITHIN GROUP (order by cats.tail_length), percentile_disc(0.5) WITHIN GROUP (order by cats.whiskers_length) FROM cats")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]*cats, 0)
	for rows.Next() {
		bk := new(cats)
		err := rows.Scan(&bk.medianTail, &bk.medianWhiskers)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range bks {
		//fmt.Println(bk.medianTail, bk.medianWhiskers)
		_, err := db.Exec("INSERT INTO cats_stat (tail_length_median, whiskers_length_median) VALUES ($1, $2)", bk.medianTail, bk.medianWhiskers)
		if err != nil {
			panic(err)
		}
	}
}
