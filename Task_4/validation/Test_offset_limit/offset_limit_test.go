package validation

import (
	"Backend_task_4/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input, expectedResult, actualResult string
var err error

//_________ТЕСТ ВАЛИДНЫХ ПАРАМЕТРОВ________________________________________________________________________________________________________________________________________________

func TestValidAttribute(t *testing.T) {

	//ФУНКЦИЯ-WhiteLimit

	input = ""
	expectedResult = ""
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "1"
	expectedResult = "limit 1"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "11"
	expectedResult = "limit 11"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "27"
	expectedResult = "limit 27"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "28"
	expectedResult = "limit 27"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "30"
	expectedResult = "limit 27"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "100"
	expectedResult = "limit 27"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	//ФУНКЦИЯ-WhiteOffset

	input = ""
	expectedResult = ""
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "0"
	expectedResult = "offset 0"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "1"
	expectedResult = "offset 1"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "11"
	expectedResult = "offset 11"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "26"
	expectedResult = "offset 26"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "27"
	expectedResult = "offset 26"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "100"
	expectedResult = "offset 26"
	actualResult, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

//_________ТЕСТ НЕВАЛИДНЫХ ПАРАМЕТРОВ____________________________________________________________________________________________________________________________________________

func TestIvalidAttribute(t *testing.T) {

	//ФУНКЦИЯ-WhiteLimit

	input = " "
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "  27"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "      "
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "DROP"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "abracadabra"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "&"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "%"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "0"
	expectedResult = "Значение не может быть нулевым"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-1"
	expectedResult = "Значение не может быть отрицательным"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-26"
	expectedResult = "Значение не может быть отрицательным"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "="
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "/"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	//ФУНКЦИЯ-WhiteOffset

	input = " "
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "  26"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "      "
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "DROP"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "abracadabra"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "&"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "%"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-1"
	expectedResult = "Значение не может быть отрицательным"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-26"
	expectedResult = "Значение не может быть отрицательным"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "="
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "/"
	expectedResult = "Задано некорректное значение"
	_, err = validation.WhiteOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}
