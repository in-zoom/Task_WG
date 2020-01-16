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

	//ФУНКЦИЯ-WhiteAttribute

	input = "name"
	expectedResult = "ORDER BY name"
	actualResult, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "tail_length"
	expectedResult = "ORDER BY tail_length"
	actualResult, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "color"
	expectedResult = "ORDER BY color"
	actualResult, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = ""
	expectedResult = ""
	expectedResult, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, expectedResult)

	//ФУНКЦИЯ-WhiteOrder

	input = "asc"
	expectedResult = "asc"
	actualResult, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = "desc"
	expectedResult = "desc"
	actualResult, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

	input = ""
	expectedResult = ""
	actualResult, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}

//_________ТЕСТ НЕВАЛИДНЫХ ПАРАМЕТРОВ____________________________________________________________________________________________________________________________________________

func TestIvalidAttribute(t *testing.T) {

	//ФУНКЦИЯ-WhiteAttribute

	input = " "
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "  color"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "      "
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "DROP"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "abracadabra"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "&"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "%"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "1"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-1"
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "="
	expectedResult = "Неверный параметр группировки"
	_, err = validation.WhiteAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	//ФУНКЦИЯ-WhiteOrder

	input = " "
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "  asc"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "      "
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "DROP"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "abracadabra"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "&"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "%"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "1"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "-1"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "="
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)

	input = "/"
	expectedResult = "Неверный параметр сортировки"
	_, err = validation.WhiteOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}
