package validation

import (
	"Backend/task_5/validation"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidCatTailLength(t *testing.T) {
	inputLength := []int{0, -1, -12, -333, -9999}
	for _, inputСurrentLength := range inputLength {
		expectedResult := "Значение не может быть равно" + " " + fmt.Sprint(inputСurrentLength)
		_, err := validation.ValidTailLength(inputСurrentLength)
		if err != nil {
			t.Error()
		}
		assert.Equal(t, expectedResult, err)
	}
}
