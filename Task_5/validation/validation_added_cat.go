package validation

import (
	"Backend/task_5/login"
	"database/sql"
	"errors"
	"fmt"
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
		lowercaseFlowerNames := prepareColor(color)
		simpleColors := strings.Split(lowercaseFlowerNames, " & ")
		var validColors = []string{"red", "black", "white"}
		for _, colorElement := range simpleColors {
			if notContains(validColors, colorElement) {
				return "", errors.New("Некорректно задан окрас кота")
			}
		}
		return lowercaseFlowerNames, nil
	}
	return "", errors.New("Не задан окрас кота")
}

func prepareColor(inputColor string) (outputLowercaseFlowerNames string) {
	colorSpaceRemoval := strings.TrimSpace(inputColor)
	lowercaseFlowerNames := strings.ToLower(colorSpaceRemoval)
	return lowercaseFlowerNames

}

func notContains(validColors []string, colorElement string) bool {
	for _, nameColor := range validColors {
		if nameColor == colorElement {
			return false
		}
	}
	return true
}

func ValidTailLength(TailLength int) (resultTailLength int, err error) {
	if TailLength <= 0 {
		return 0, errors.New("Значение не может быть равно" + " " + fmt.Sprint(TailLength))
	}
	return TailLength, nil
}
