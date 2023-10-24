package colors

import "testing"

// Given a valid Hex color
// Then it should return a RGB color and no error
func TestConvertHexToRGBIsValid(t *testing.T) {
	result, err := ConvertHexToRGB([]string{"#FFFFFF"})
	expected := "255, 255, 255"

	if err != nil {
		t.Fatalf("Should not have any error, got %v", err)
	}
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

// Given an invalid Hex color
// Then it should return an error
func TestHexValueIsNotValid(t *testing.T) {
	valuesToTest := []string{"#FFFFFZ", "FFFFFF", "", "#FFF"}

	for _, val := range valuesToTest {
		_, err := ConvertHexToRGB([]string{val})

		if err == nil {
			t.Fatalf("Should have an error, got nil")
		}
	}
}
