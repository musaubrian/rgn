package utils

import "testing"

func TestTruncate(t *testing.T) {
	text := "This is a very long string"
	maxLen := 10
	plusEllipsis := maxLen + 3
	res := Truncate(text, maxLen)
	if len(res) > plusEllipsis {
		t.Error(res)
		t.Errorf("Expected %d, got %d", plusEllipsis, len(res))
	}
}
