package Mean

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Mean(db *sql.DB, columName string) float32 {
	//rows, err := db.Query("SELECT SUM (col) / COUNT(cats.name) FROM cats WHERE col = $1", columName)
	rows, err := db.Query(fmt.Sprintf("SELECT SUM (%s) / COUNT(cats.name) FROM cats", columName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ResultMean float32
	for rows.Next() {
		err := rows.Scan(&ResultMean)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ResultMean
}
