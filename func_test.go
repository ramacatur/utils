package utils

import (
	"testing"
)

func TestPadLeft(t *testing.T) {
	result := PadLeft("2", "0", 4)

	if result != "0002" {
		t.Fatalf("Expect %q, actual %q", "0002", result)
	}
}

func TestPadRight(t *testing.T) {
	result := PadRight("2", "0", 2)

	if result != "20" {
		t.Fatalf("Expect %q, actual %q", "20", result)
	}
}
