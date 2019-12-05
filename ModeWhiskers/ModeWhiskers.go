package ModeWhiskers

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
)

type cats struct {
			modeWhiskers float32
}

func ModeWhiskers(db *sql.DB) {
	rows, err := db.Query("SELECT cats.whiskers_length FROM cats GROUP BY cats.whiskers_length HAVING COUNT(*) >= all (SELECT COUNT(*) FROM cats GROUP BY cats.whiskers_length)")
	//rows, err := db.Query("SELECT mode() WITHIN GROUP (ORDER BY cats.tail_length, cats.whiskers_length) AS modal_value, modal_value2 FROM cats")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := make([]*cats, 0)
	for rows.Next() {
		bk := new(cats)
		err := rows.Scan(&bk.modeWhiskers)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	modeWhiskers := make([]float32, 0)
	for _, bk := range bks {
		modeWhiskers = append(modeWhiskers, bk.modeWhiskers)
	}
	ins := "INSERT INTO cats_stat (whiskers_length_mode) VALUES ($1)"
	//fmt.Println(modeWhiskers)
	_, err = db.Exec(ins, pq.Array(modeWhiskers))
	if err != nil {
		panic(err)
	}

}
