package gh

import (
	"context"
	"testing"
)

func TestCorrectDets(t *testing.T) {
	ePath := "../../.env"
	ctx := context.Background()
	_, err := Auth(ePath, ctx)
	if err != nil {
		t.Errorf("Did not expect an error, got %v", err)
	}
}
