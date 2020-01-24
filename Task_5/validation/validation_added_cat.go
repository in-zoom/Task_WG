package validation

import (
	"Backend/task_5/login"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"strings"
)

func ValidateName(addedNameCat string) (nameCat string, err error) {
	if addedNameCat != "" {
		db := login.Init()
		var rows *sql.Rows
		var err error
		addNameCat := strings.Title(addedNameCat)
		query := "SELECT name FROM cats WHERE name = " + " " + "'" + addNameCat + "'"
		rows, err = db.Query(query)

		if err != nil {
			return "", err
		}
		defer rows.Close()
		var nameCatFromExisting string
		for rows.Next() {
			if err = rows.Scan(&nameCatFromExisting); err != nil {
				return "", err
			}
		}
		if err = rows.Err(); err != nil {
			return "", err
		}
		if nameCatFromExisting == "" {
			return addNameCat, nil
		} else {
			return "", errors.New("Кот с именем" + " " + nameCatFromExisting + " " + "уже существует")
		}

	}
	return "", errors.New("Введите имя")
}