package gh

import (
	"context"

	"github.com/google/go-github/github"
)

// Print the gitignore contents to stdout
// Can be redirected to a file(.gitignore)
// Incase one already has a repo and wants to add the .gitignore
func QuickIgnore(c *github.Client, ctx context.Context, lang string) (string, error) {
	ignoreContents, err := GetGitignore(c, ctx, lang)
	if err != nil {
		return ignoreContents, err
	}
	return ignoreContents, err
}
