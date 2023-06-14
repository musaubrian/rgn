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

// func PrettyAlignIssues(issues []*github.Issue)[][]string {
//     var finalIssueLists [][]string
//     var singleIssueList []string
//     for _, i := range issues {
//         singleIssueList = append(singleIssueList, *i.Title, i.)
//
//     }
// }
