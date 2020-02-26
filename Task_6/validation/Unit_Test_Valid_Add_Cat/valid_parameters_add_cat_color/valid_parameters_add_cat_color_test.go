package validation

import (
	"Backend_Task_5/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidCatColorBlack(t *testing.T){
	inputColor := []string {"Black", "    Black", "      Black          "}
	expectedResult := "black" 
   for _, currentColor := range inputColor {
	   actualResult, err := validation.ValidateColor(currentColor)
   if err != nil {
	   t.Error()
   }
   assert.Equal(t, expectedResult, actualResult)
   
 }
}

func TestValidCatColorWhite(t *testing.T){
   inputColor := []string {"white", "White", "    White", "      white          "}
   expectedResult := "white" 
  for _, currentColor := range inputColor {
	  actualResult, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, expectedResult, actualResult)
  
}
}

func TestValidCatColorRed(t *testing.T){
   inputColor := []string {"red", "Red", "    Red", "      red          "}
   expectedResult := "red" 
  for _, currentColor := range inputColor {
	  actualResult, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, expectedResult, actualResult)
  
}
}

func TestValidCatColorBlackWhite(t *testing.T){
   inputColor := []string {"black & white", "Black & White", "    Black & White", "      black & white          "}
   expectedResult := "black & white" 
  for _, currentColor := range inputColor {
	  actualResult, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, expectedResult, actualResult)
  
}
}

func TestValidCatColorRedWhite(t *testing.T){
   inputColor := []string {"red & white", "Red & White", "    Red & White", "      red & white          "}
   expectedResult := "red & white" 
  for _, currentColor := range inputColor {
	  actualResult, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, expectedResult, actualResult)
  
}
}

func TestValidCatColorRedBlackWhite(t *testing.T){
   inputColor := []string {"red & black & white", "Red & Black & White", "    Red & Black & White", "      red & black & white         "}
   expectedResult := "red & black & white" 
  for _, currentColor := range inputColor {
	  actualResult, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, expectedResult, actualResult)
  
}
}