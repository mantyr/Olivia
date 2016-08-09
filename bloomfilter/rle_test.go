package olilib

import (
	"testing"
)

func TestEncode(t *testing.T) {
	expectedReturn := "A3BZ5T3"
	retVal := Encode("AAABZZZZZTTT")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestDecode(t *testing.T) {
	expectedReturn := "AAABZZZZZTTT"
	retVal := Decode("A3BZ5T3")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestDecodeLongString(t *testing.T) {
	expectedReturn := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	retVal := Decode("a30")

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestWriteOutputOver9(t *testing.T) {
	expectedReturn := "a63"
	retVal := writeOutput("", 'a', 63)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestWriteOutputOver9WithExtra(t *testing.T) {
	expectedReturn := "a68"
	retVal := writeOutput("", 'a', 68)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}

func TestEncodeDecode(t *testing.T) {
	expectedReturn := "AAABBZZZZZTTT"
	retVal := Encode(expectedReturn)
	retVal = Decode(retVal)

	if expectedReturn != retVal {
		t.Errorf("Expected %v, got %v", expectedReturn, retVal)
	}
}
