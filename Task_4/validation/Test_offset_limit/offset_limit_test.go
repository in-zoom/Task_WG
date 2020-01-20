package validation

import (
	"Backend_task_4/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

var input, expectedResult, actualResult string
var err error

func TestValidAttribute(t *testing.T) {
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

	input = "100"
	expectedResult = "limit 27"
	actualResult, err = validation.WhiteLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)

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

	input = "26"
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
func TestIvalidAttribute(t *testing.T) {
	imput := []string{"-1", "0", " ", "DROP", "&"}
	for _, imputCurrentItem := range imput {
		expectedResult = "Значение не может быть отрицательным, Значение не может быть нулевым, Задано некорректное значение"
		_, err = validation.WhiteLimit(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)

		expectedResult = "Значение не может быть отрицательным, Задано некорректное значение"
		_, err = validation.WhiteOffset(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
