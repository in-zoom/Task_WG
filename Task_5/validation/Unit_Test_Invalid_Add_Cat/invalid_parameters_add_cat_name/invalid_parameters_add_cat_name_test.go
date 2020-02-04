package validation

import (
	"Backend/task_5/validation"
	"testing"
    "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInvalidIfNameCatExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := []string{"Marik", "marik", "  marik", "   marik       "}
	expectedResult := "Кот с именем Marik уже существует"
	for _, currentName := range input {
		rows := sqlmock.NewRows([]string{"name"}).AddRow("Marik")
		mock.ExpectQuery("SELECT name FROM cats WHERE name = " + " " + "'" + "Marik" + "'").WillReturnRows(rows)
		_, err := validation.ValidateName(currentName, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)
	}
}

func TestInvalidIfNameCatThereAreCharactersAndNumbers(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputInvalidName := []string{"Marik!", "marik1", "  mar%ik ", "   & marik       ", " /marik", " Marik. ", "5", "?", "+"}
	expectedResult := "Имя не может содержать цифры, знаки пунктуации, символы"
	for _, currentInvalidName := range inputInvalidName {
		_, err := validation.ValidateName(currentInvalidName, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}

func TestInvalidIfEmptyLine(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	inputInvalidName := []string{"", " ", "      "}
	expectedResult := "Введите имя"
	for _, currentInvalidNameEmptyLine := range inputInvalidName {
		_, err := validation.ValidateName(currentInvalidNameEmptyLine, db)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, err, expectedResult)

	}
}
