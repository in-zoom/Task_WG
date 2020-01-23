package validation

import (
	"Backend/task_4/login"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"strings"
)

func NameValid(addedNameCat string) (nameCat string, err error) {
	if addedNameCat != "" {
		db := login.Init()
		var rows *sql.Rows
		var err error
		query := "SELECT name FROM cats"
		rows, err = db.Query(query)

		if err != nil {
			return "", err
		}
		defer rows.Close()
		var nameCatFromExisting, addNameCat string
		for rows.Next() {
			if err = rows.Scan(&nameCatFromExisting); err != nil {
				return "", err
			}
			addNameCat = strings.Title(addedNameCat)
			if addNameCat == nameCatFromExisting {
				return "", errors.New("Кот с указанным именем существует")
			}
		}
		if err = rows.Err(); err != nil {
			return "", err
		}
		return addNameCat, nil
	}
	return "", errors.New("Введите имя")
}
