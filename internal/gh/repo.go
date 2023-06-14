package gh

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateEmptyRepo(c *github.Client, ctx context.Context) (*github.Repository, error) {
	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return nil, custom.GetGHUserErr(err)
	}

	// Repo name must not be empty
	rName, err := utils.ReadInput("Repo name:")
	if err != nil {
		return nil, err
	}
	rDesc, _ := utils.ReadInput("Repo description:")
	visibility, _ := utils.ReadInput("Make private [Y]/n:")
	private := true
	defBranch := "main"
	if visibility == "n" {
		private = false
	}

	rOpts := github.Repository{
		Owner:         &github.User{Login: github.String(*u.Name)},
		Name:          &rName,
		Description:   &rDesc,
		DefaultBranch: &defBranch,
		MasterBranch:  &defBranch,
		Private:       &private,
	}

	repo, _, err := c.Repositories.Create(ctx, "", &rOpts)
	if err != nil {
		return repo, custom.CreateRepoErr(err)
	}
	return repo, nil
}

func CreateRepoWithReadme(c *github.Client, ctx context.Context) (*github.Repository, error) {
	r, err := CreateEmptyRepo(c, ctx)
	if err != nil {
		return nil, err
	}
	err = CreateReadme(c, ctx, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreateRepoWithGitignore(c *github.Client, ctx context.Context) (*github.Repository, error) {
	r, err := CreateEmptyRepo(c, ctx)
	if err != nil {
		return nil, err
	}

	lang, err := utils.ReadInput("Language to setup gitignore for:")
	if err != nil {
		return nil, err
	}
	// Capitalize
	lang = cases.Title(language.English, cases.NoLower).String(lang)

	err = CreateGitignore(c, ctx, r, lang)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreateRepoWithBoth(c *github.Client, ctx context.Context) (*github.Repository, error) {
	r, err := CreateEmptyRepo(c, ctx)
	if err != nil {
		return nil, err
	}
	err = CreateReadme(c, ctx, r)
	if err != nil {
		return nil, err
	}

	lang, err := utils.ReadInput("Language to setup gitignore for:")
	if err != nil {
		return nil, err
	}
	// Capitalize
	lang = cases.Title(language.English, cases.NoLower).String(lang)

	err = CreateGitignore(c, ctx, r, lang)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreateReadme(c *github.Client, ctx context.Context, r *github.Repository) error {
	commitMsg := "chore: create README"
	fContent := "# " + *r.Name
	currentTime := time.Now()

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return custom.GetGHUserErr(err)
	}

	a := github.CommitAuthor{
		Date:  &currentTime,
		Name:  u.Name,
		Email: u.Email,
		Login: u.Login,
	}

	fOpts := github.RepositoryContentFileOptions{
		Message:   &commitMsg,
		Content:   []byte(fContent),
		Author:    &a,
		Committer: &a,
	}
	_, _, err = c.Repositories.CreateFile(ctx, *a.Login, *r.Name, "README.md", &fOpts)
	if err != nil {
		return custom.FileCreationErr("README.md", err)
	}
	return nil
}

func CreateGitignore(c *github.Client, ctx context.Context, r *github.Repository, lang string) error {
	commitMsg := "chore: create .gitignore"
	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return custom.GetGHUserErr(err)
	}
	gitIgnoreContent, err := GetGitignore(c, ctx, lang)
	if err != nil {
		return err
	}
	currentTime := time.Now()

	a := github.CommitAuthor{
		Date:  &currentTime,
		Name:  u.Name,
		Email: u.Email,
		Login: u.Login,
	}

	fOpts := github.RepositoryContentFileOptions{
		Message:   &commitMsg,
		Content:   []byte(gitIgnoreContent),
		Author:    &a,
		Committer: &a,
	}
	_, _, err = c.Repositories.CreateFile(ctx, *a.Login, *r.Name, ".gitignore", &fOpts)
	if err != nil {
		return custom.FileCreationErr(".gitignore", err)
	}
	return nil
}

func GetGitignore(c *github.Client, ctx context.Context, lang string) (string, error) {
	gitIgnore, _, err := c.Gitignores.Get(ctx, lang)
	if err != nil {
		return "", custom.GetGitignoreErr(err)
	}
	return gitIgnore.String(), nil
}
