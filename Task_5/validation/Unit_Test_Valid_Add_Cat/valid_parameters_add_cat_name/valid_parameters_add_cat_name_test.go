package validation

import (
	"github.com/DATA-DOG/go-sqlmock"
	"Backend/task_5/validation"
	"testing"
)

func TestValidNameCat(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	input := []string{
		"Marik",
		"marik",
		"  marik",
		"   marik       ",
		"MariK",
		" MARIK     ",
		"    maRiK   "}

	for _, currentName := range input {
		rows := sqlmock.NewRows([]string{"name"})
		mock.ExpectQuery("SELECT name FROM cats WHERE name = " + " " + "'" + "Marik" + "'").WillReturnRows(rows)
		_, err := validation.ValidateName(currentName, db)
		if err != nil {
			t.Error()
		}
	}
}
