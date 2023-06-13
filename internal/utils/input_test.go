package utils

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := "Hello, World!\n"
	stdin := bytes.NewBufferString(input)
	oldStdin := os.Stdin
	os.Stdin = createFile(stdin)
	defer func() {
		os.Stdin = oldStdin
		os.Stdin.Close()
	}()

	prompt := "Enter some input:"
	expected := "Hello, World!"

	result, err := ReadInput(prompt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}

	// Test empty input
	emptyInput := "\n"
	stdin = bytes.NewBufferString(emptyInput)
	os.Stdin = createFile(stdin)

	expectedErr := errors.New("No input entered")
	result, err = ReadInput(prompt)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
	}
}

// Helper function to create *os.File from io.Reader
func createFile(r io.Reader) *os.File {
	f, _ := os.CreateTemp("", "")
	io.Copy(f, r)
	f.Seek(0, 0)
	return f
}
