package strings

import "testing"

func TestCountIsSuccessWithWesternCharset(t *testing.T) {
	expected := "5"
	actual, _ := Count([]string{"hello"})
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCountIsSucessWithJapaneseCharset(t *testing.T) {
	expected := "5"
	actual, _ := Count([]string{"こんにちは"})
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
