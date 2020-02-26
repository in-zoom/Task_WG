package handlers

import (
	"Backend_task_6/login"
	_ "github.com/lib/pq"
	"database/sql"
	
)

var err error

func сatslist(attribute string, order string, offset string, limit string) ([]cat, error) {
	db := login.Connect()
	defer db.Close()
	var rows *sql.Rows
	query := "SELECT * FROM cats" + " " + attribute + " " + order + " " + offset + " " + limit
	rows, err = db.Query(query)

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

func addNewCat(nameCat, colorCat string, TailLengthCat, WhiskersLength int) (err error) {
	db := login.Connect()
	ins := "INSERT INTO cats (name, color, tail_length, whiskers_length) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(ins, nameCat, colorCat, TailLengthCat, WhiskersLength)
	if err != nil {
		return err
	}
	return nil
}
