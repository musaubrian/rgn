package utils

import (
	"strings"
	"testing"
	"time"
)

func TestTimeInDays(t *testing.T) {
	dur := 24 * time.Hour
	res := TimeInDays(dur)
	if !strings.Contains(res, "1d") {
		t.Errorf("Expected 1d, got %s", res)
	}
}
