package validation

import (
	"Backend/Task_4/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestСolumnName(t *testing.T) {
	input := "name"
	expectedResult := "ORDER BY name"
	actualResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestСolumnTailLength(t *testing.T) {
	input := "tail_length"
	expectedResult := "ORDER BY tail_length"
	actualResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestСolumnColor(t *testing.T) {
	input := "color"
	expectedResult := "ORDER BY color"
	actualResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestСolumnEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	expectedResult, err := validation.ValidateAttribute(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, expectedResult)
}
func TestSortingAscending(t *testing.T) {
	input := "asc"
	expectedResult := "asc"
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestSortingDescending(t *testing.T) {
	input := "desc"
	expectedResult := "desc"
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestSortingEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateOrder(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfLimitEqualEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfLimitEqualOne(t *testing.T) {
	input := "1"
	expectedResult := "limit 1"
	actualResult, err := validation.ValidateLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfLimitEqualOneHundred(t *testing.T) {
	input := "100"
	expectedResult := "limit 27"
	actualResult, err := validation.ValidateLimit(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfOffsetEqualEmpty(t *testing.T) {
	input := ""
	expectedResult := ""
	actualResult, err := validation.ValidateOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfOffsetEqualZero(t *testing.T) {
	input := "0"
	expectedResult := "offset 0"
	actualResult, err := validation.ValidateOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfOffsetEqualTwentySix(t *testing.T) {
	input := "26"
	expectedResult := "offset 26"
	actualResult, err := validation.ValidateOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
func TestIfOffsetEqualOneHundred(t *testing.T) {
	input := "100"
	expectedResult := "offset 26"
	actualResult, err := validation.ValidateOffset(input)
	if err != nil {
		t.Error()
	}
	assert.Equal(t, expectedResult, actualResult)
}
