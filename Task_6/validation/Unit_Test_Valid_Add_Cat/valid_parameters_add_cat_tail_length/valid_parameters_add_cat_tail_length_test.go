package validation

import (
	"Backend_Task_5/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidCatTailLength(t *testing.T){
	inputLength := []int { 1, 12, 333, 9999}
	for _, inputСurrentLength := range inputLength{
	expectedResult := inputСurrentLength
   actualResult, err := validation.ValidTailLength(inputСurrentLength)
   if err != nil {
	   t.Error()
   }
   assert.Equal(t, expectedResult, actualResult)
}
}