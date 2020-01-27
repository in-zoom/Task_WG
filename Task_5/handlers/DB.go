package handlers

import (
	"Backend/task_5/login"
	"database/sql"
	_ "github.com/lib/pq"
)

var err error

func сatslist(attribute string, order string, offset string, limit string) ([]cat, error) {
	db := login.Init()
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

func addNewCat(nameCat, colorCat string) (err error) {
	db := login.Init()
	ins := "INSERT INTO cats (name, color) VALUES ($1, $2)"
	_, err = db.Exec(ins, nameCat, colorCat)
	if err != nil {
		return err
	}
	return nil
}
