package colors

import "testing"

// Given a valid RGB color
// Then it should return the corresponding Hex color and no error
func TestConvertRGBToHexIsValid(t *testing.T) {
	result, err := ConvertRGBToHex([]string{"255", "255", "255"})
	expected := "#FFFFFF"

	if err != nil {
		t.Fatalf("Should not have any error, got %v", err)
	}
	if result != expected {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

// Given an invalid RGB color
// Then it should return an error
func TestConvertRGBToHexIsNotValid(t *testing.T) {
	values := [][]string{
		{"255", "255", "255", "255"},
		{"255", "255"},
		{"-4", "255", "255"},
	}

	for _, val := range values {
		_, err := ConvertRGBToHex(val)
		if err == nil {
			t.Fatalf("Should have an error, got nil for: %v %T", val, val)
		}
	}
}
