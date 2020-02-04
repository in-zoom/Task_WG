package validation

import (
	"Backend_Task_5/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidCatWhiskersLength(t *testing.T) {
	inputWhiskersLength := []int{1, 12, 333, 9999}
	for _, inputСurrentWhiskersLength := range inputWhiskersLength {
		expectedResult := inputСurrentWhiskersLength
		actualResult, err := validation.ValidWhiskersLength(inputСurrentWhiskersLength)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, actualResult)
	}
}
