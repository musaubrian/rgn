package gh

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/custom"
)

func GetIssuesAssigned(c *github.Client, ctx context.Context) ([]*github.Issue, error) {
	iOpts := github.IssueListOptions{
		Filter:    "assigned",
		State:     "open",
		Direction: "",
	}
	i, _, err := c.Issues.List(ctx, true, &iOpts)
	if err != nil {
		return nil, custom.GetIssuesErr(err)
	}
	return i, nil
}
