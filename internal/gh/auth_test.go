package gh

import (
	"context"
	"testing"

	"github.com/musaubrian/rgn/internal/utils"
)

func TestCorrectDets(t *testing.T) {
	ePath, _ := utils.GetEnvLoc()
	ctx := context.Background()
	_, err := Auth(ePath, ctx)
	if err != nil {
		t.Errorf("Did not expect an error, got %v", err)
	}
}
