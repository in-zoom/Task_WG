package Mode

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func Mode(db *sql.DB, columName string) []float32 {
	rows, err := db.Query(fmt.Sprintf("SELECT %s FROM cats GROUP BY %s HAVING COUNT(*) >= all (SELECT COUNT(*) FROM cats GROUP BY %s)", columName, columName, columName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var result float32
	ResultMode := make([]float32, 0)
	for rows.Next() {
		err := rows.Scan(&result)
		ResultMode = append(ResultMode, result)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ResultMode
}
