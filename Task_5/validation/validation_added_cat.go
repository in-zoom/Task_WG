package validation

import (
	"Backend/task_5/login"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"regexp"
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
	validColor := make([]string, 0)
	if color != "" {
		sentenceWord := strings.Split(color, " ")
		for _, colorName := range sentenceWord {
			validateListColor := []string{"white", "black", "red", "&"}
			for _, currentСolor := range validateListColor {
				if colorName == currentСolor {
					validColor = append(validColor, colorName)
				}
			}
		}
		if len(validColor) == 0 {
			return "", errors.New("Некорректно задан окрас кота")
		}
		for _, currentNameColor := range validColor {
			if resultColor == "" {
				resultColor += currentNameColor
			} else {
				resultColor += " " + currentNameColor
			}
		}
		pattern := `(^[A-Za-z]+$)|(^[A-Za-z]+[\ t \ n \ v \ f \ r]+&+[\ t \ n \ v \ f \ r]+[A-Za-z]+$)|(^[A-Za-z]+[\ t \ n \ v \ f \ r]+&+[\ t \ n \ v \ f \ r]+[A-Za-z]+[\ t \ n \ v \ f \ r]+&+[\ t \ n \ v \ f \ r]+[A-Za-z]+$)`
		matched, err := regexp.Match(pattern, []byte(resultColor))
		if matched == false || err != nil {
			return "", errors.New("Некорректно задан окрас кота")
		} else {
			return resultColor, nil
		}
	}
	return "", errors.New("Не задан окрас кота")
}
