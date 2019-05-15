package db

import (
	"testing"
)

func TestValidKey(t *testing.T) {
	key := "abcXYZ123"
	err := validateKey(key)

	if err != nil {
		t.Fatalf("Key %s should be valid", key)
	}
}

func TestInvalidKey(t *testing.T) {
	key := "abcXYZ-123"
	err := validateKey(key)

	if err == nil {
		t.Fatalf("Key %s should be invalid", key)
	}
}
