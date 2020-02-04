package validation

import (
	//"Backend/task_5/login" // раскомментировать для нормальной работы
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

func ValidateName(addedNameCat string, db *sql.DB) (nameCat string, err error) { // для тестирования
//func ValidateName(addedNameCat string) (nameCat string, err error) { // раскомментировать для нормальной работы
	if addedNameCat != "" {
		//db := login.Init() // раскомментировать для нормальной работы
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
	lowercaseFlowerNames := prepareColor(color)
	if color != "" {
		var validColors = []string{
			"black",
			"white",
			"black & white",
			"red",
			"red & white",
			"red & black & white"}

		for _, currentСolor := range validColors {
			if currentСolor == lowercaseFlowerNames {
				return lowercaseFlowerNames, nil
			}
		}
		return "", errors.New("Некорректно задан окрас кота")
	}
	return "", errors.New("Не задан окрас кота")
}

func prepareColor(inputColor string) (outputLowercaseFlowerNames string) {
	colorSpaceRemoval := strings.TrimSpace(inputColor)
	lowercaseFlowerNames := strings.ToLower(colorSpaceRemoval)
	return lowercaseFlowerNames

}

func ValidTailLength(TailLength int) (resultTailLength int, err error) {
	if TailLength <= 0 {
		return 0, errors.New("Значение не может быть равно" + " " + fmt.Sprint(TailLength))
	}
	return TailLength, nil
}
func ValidWhiskersLength(WhiskersLength int) (resultWhiskersLength int, err error) {
	if WhiskersLength <= 0 {
		return 0, errors.New("Значение не может быть равно" + " " + fmt.Sprint(WhiskersLength))
	}
	return WhiskersLength, nil
}
