package utils

import (
	"os"
	"testing"
)

func TestReadEnv(t *testing.T) {
	val, err := ReadEnv("../../.env")
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Expected .env to exist")
		}
	}

	if len(val) != 2 {
		t.Errorf("Expected 2 values, got %d", len(val))
	}
}

func TestCreateEnv(t *testing.T) {
	// pass
}
