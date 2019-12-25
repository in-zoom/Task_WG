package handlers

import (
	"Backend_task_4/login"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func сatslist(param string) ([]cat, error) {
	db := login.Init()
	var rows *sql.Rows
	var err error
	ins := "SELECT * FROM cats" + " " + param
	rows, err = db.Query(fmt.Sprintf("%s", ins))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]cat, 0)
	var cat cat
	for rows.Next() {
		if err = rows.Scan(&cat.Name, &cat.Color, &cat.TailLength, &cat.WhiskersLength); err != nil {
			return nil, err
		}
		list = append(list, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
