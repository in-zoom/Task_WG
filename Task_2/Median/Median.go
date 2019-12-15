package Median

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Median(db *sql.DB, columName string) float32 {
	rows, err := db.Query(fmt.Sprintf("SELECT percentile_disc(0.5) WITHIN GROUP (order by %s) FROM cats", columName))
	//rows, err := db.Query("SELECT percentile_disc(0.5) WITHIN GROUP (order by &1)", columName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ResultMedian float32
	for rows.Next() {
		err := rows.Scan(&ResultMedian)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ResultMedian
}
