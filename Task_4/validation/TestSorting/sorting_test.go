package validation

import (
	"Backend_task_4/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input, expectedResult, actualResult string
var err error

func TestValidAttribute(t *testing.T) {
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
func TestIvalidAttribute(t *testing.T) {
	imput := []string{" ", "DROP", "&", "1"}
	for _, imputCurrentItem := range imput {
		expectedResult = "Неверный параметр группировки"
		_, err = validation.WhiteAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
		expectedResult = "Неверный параметр сортировки"
		_, err = validation.WhiteOrder(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
