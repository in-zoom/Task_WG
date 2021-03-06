package validation

import (
	"Backend/Task_4/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameСolumnSigns(t *testing.T) {
	imput := []string{"&", "%", "*", "="}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр группировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestNameСolumnStrings(t *testing.T) {
	imput := []string{" ", "DROP", "5"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр группировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestSortingSigns(t *testing.T) {
	imput := []string{"&", "%", "*", "="}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр сортировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
func TestSortingStrings(t *testing.T) {
	imput := []string{" ", "DROP", "5"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Неверный параметр сортировки"
		_, err := validation.ValidateAttribute(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}

func TestNegativeValueLimit(t *testing.T) {
	imput := "-1"
	expectedResult := "Значение не может быть отрицательным"
	_, err := validation.ValidateLimit(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestZeroValueLimit(t *testing.T) {
	imput := "0"
	expectedResult := "Значение не может быть нулевым"
	_, err := validation.ValidateLimit(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestIncorrectStringsLimit(t *testing.T) {
	imput := []string{" ", "DROP", "&"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Задано некорректное значение"
		_, err := validation.ValidateLimit(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}

func TestNegativeValueOffset(t *testing.T) {
	imput := "-1"
	expectedResult := "Значение не может быть отрицательным"
	_, err := validation.ValidateOffset(imput)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, err)
}

func TestIncorrectStringsOffset(t *testing.T) {
	imput := []string{" ", "DROP", "&"}
	for _, imputCurrentItem := range imput {
		expectedResult := "Задано некорректное значение"
		_, err := validation.ValidateOffset(imputCurrentItem)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
