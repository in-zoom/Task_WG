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

func ValidateColor(color string) (resultColor string, err error) {
	if color != "" {
		colorSpaceRemoval := strings.TrimSpace(color)
		lowercaseFlowerNames := strings.ToLower(colorSpaceRemoval)
		var validColors = []string{"red", "black", "white"}
		simpleColors := strings.Split(lowercaseFlowerNames, " & ")
		for _, colorElement := range simpleColors {
			if notContains(validColors, colorElement) {
				return "", errors.New("Некорректно задан окрас кота")
			}
		}
		return color, nil
	}
	return "", errors.New("Не задан окрас кота")
}

func notContains(validColors []string, colorElement string) bool {
	for _, nameColor := range validColors {
		if nameColor == colorElement {
			return false
		}
	}
	return true
}
