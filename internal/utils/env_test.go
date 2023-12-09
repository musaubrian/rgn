package utils

import (
	"os"
	"testing"
)

func TestReadEnv(t *testing.T) {
	envPath, _ := GetEnvLoc()
	val, err := ReadEnv(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Expected .env to exist")
		}
	}

	if len(val) != 2 {
		t.Errorf("Expected 2 values, got %d", len(val))
	}
}

func TestGetEnvLoc(t *testing.T) {
	_, err := GetEnvLoc()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
