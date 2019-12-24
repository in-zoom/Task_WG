package handlers

import (
	"Backend/Task_4/login"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Catslist(param1 string, param2 string, param3 string, param4 string) ([]Cats, error) {
	db := login.Init()
	var rows *sql.Rows
	var err error
	param := param1 + " " + param2
	rows, err = db.Query(fmt.Sprintf("SELECT * FROM cats ORDER BY %s offset %s limit %s", param, param3, param4))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]Cats, 0)
	var cat Cats
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
