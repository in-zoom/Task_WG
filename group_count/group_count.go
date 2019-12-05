package group_count

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
)

type cats struct {
	color         string
	quantityColor float32
}

func GroupCount(db *sql.DB) {
	rows, err := db.Query("SELECT color, COUNT(color)  FROM cats GROUP BY color ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]*cats, 0)
	for rows.Next() {
		bk := new(cats)
		err := rows.Scan(&bk.color, &bk.quantityColor)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, bk := range bks {
		_, err := db.Exec("INSERT INTO cat_colors_info (color, count) VALUES ($1, $2)", bk.color, bk.quantityColor)
		if err != nil {
			panic(err)
		}
	}
}
