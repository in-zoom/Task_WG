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
	rows, err := db.Query("SELECT color, COUNT(color)  FROM cats GROUP BY color")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	groupCount := new(cats)
	for rows.Next() {
		err := rows.Scan(&groupCount.color, &groupCount.quantityColor)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	ins := "INSERT INTO cat_colors_info (color, count) VALUES ($1, $2)"
	_, errr := db.Exec(ins, groupCount.color, groupCount.quantityColor)
	if errr != nil {
		panic(errr)
	}
}
