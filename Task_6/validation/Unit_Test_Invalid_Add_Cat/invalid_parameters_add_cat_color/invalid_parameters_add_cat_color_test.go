package validation

import (
	"Backend_task_5/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidArbitraryWord(t *testing.T){
  inputColor := []string {"abracadabra", "    Abracadabra", "      abracadabra          "}
  expectedResult := "Некорректно задан окрас кота"
   for _, currentColor := range inputColor {
	   _, err := validation.ValidateColor(currentColor)
   if err != nil {
	   t.Error()
   }
   assert.Equal(t, err, expectedResult)
   
 }
}

func TestInvalidColorCat(t *testing.T){
   inputColor := []string {"yellow", "Yellow", "    Yellow", "      yellow          "}
   expectedResult := "Некорректно задан окрас кота"
  for _, currentColor := range inputColor {
	  _, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, err, expectedResult)
  
}
}

func TestInvalidSymbol(t *testing.T){
   inputColor := []string {"!", "&", "    *", "      %          "}
   expectedResult := "Некорректно задан окрас кота"
  for _, currentColor := range inputColor {
	  _, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, err, expectedResult)
  
}
}

func TestInvalidEmptyLines(t *testing.T){
   inputColor := []string {"", " ", "        "}
   expectedResult := "Не задан окрас кота"
  for _, currentColor := range inputColor {
	  _, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, err, expectedResult)
  
}
}

func TestInvalidArbitraryWordAndValidColor(t *testing.T){
   inputColor := []string {"abracadabra & white", "Red & Abracadabra", "    abracadabra & white", "      Red & Abracadabra          "}
   expectedResult := "Некорректно задан окрас кота"
  for _, currentColor := range inputColor {
	  _, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, err, expectedResult)
  
}
}

func TestInvalidArbitraryWordAndValidColorOfThreeElements(t *testing.T){
   inputColor := []string {"red & yellow & white", "Yellow & Black & White", "    Red & Black & Yellow", "      yellow & black & white         "}
   expectedResult := "Некорректно задан окрас кота"
  for _, currentColor := range inputColor {
	  _, err := validation.ValidateColor(currentColor)
  if err != nil {
	  t.Error()
  }
  assert.Equal(t, err, expectedResult)
  
}
}