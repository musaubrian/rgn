package gh

import (
	"context"

	"github.com/google/go-github/github"
)

// Do a form of token validity check
func Ping(c *github.Client, ctx context.Context) error {
	_, _, err := c.Users.Get(ctx, "")

	return err
}
