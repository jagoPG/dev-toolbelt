package encoding

import "testing"

// Given a string
// Return a base64 encoded string
func TestEncodeBase64(t *testing.T) {
	expected := "Zm9v"

	result, err := EncodeBase64([]string{"foo"})
	if err != nil {
		t.Errorf("Unexpected error returned: %s", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestDecodeBase64(t *testing.T) {
	expected := "foo"

	result, err := DecodeBase64([]string{"Zm9v"})
	if err != nil {
		t.Errorf("Unexpected error returned: %s", err)
	}
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
