package arguments

import (
	"testing"
)

// Given some arguments and a valid expected amount
// Then it should return true and no error
func TestValidAmountOfArguments(t *testing.T) {
	result, err := ValidateArgs([]string{"255", "255", "255"}, 3)

	if !result {
		t.Fatalf("Expected true, got %v", result)
	}
	if err != nil {
		t.Fatalf("Should not have any error, got %v", err)
	}
}

// Given some arguments and an invalid expected amount
// Then it should return false and an error
func TestInvalidAmountOfArguments(t *testing.T) {
	result, err := ValidateArgs([]string{"255", "255"}, 3)

	if result || err == nil {
		t.Fatalf("Expected false and error, got %v, %v", result, err)
	}
}
