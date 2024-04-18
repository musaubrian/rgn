package gh

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreateEmptyRepo(c *github.Client, ctx context.Context) (*github.Repository, error) {
	var makePrivate bool

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return nil, custom.Error(custom.ErrMsg["getGhUserErr"], err)
	}

	// Repo name must not be empty
	rName, err := utils.ReadInput("Repo name")
	if err != nil {
		return nil, err
	}
	rDesc, _ := utils.ReadInput("Repo description")

	err = huh.NewConfirm().
		Title("Make private").
		Value(&makePrivate).
		Run()

	if err != nil {
		log.Fatal(err)
	}
	defBranch := "main"

	rOpts := github.Repository{
		Owner:         &github.User{Login: github.String(*u.Name)},
		Name:          &rName,
		Description:   &rDesc,
		DefaultBranch: &defBranch,
		MasterBranch:  &defBranch,
		Private:       &makePrivate,
	}

	repo, _, err := c.Repositories.Create(ctx, "", &rOpts)
	if err != nil {
		return repo, custom.Error(custom.ErrMsg["repoExistsErr"], err)
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
	var lang string
	r, err := CreateEmptyRepo(c, ctx)
	if err != nil {
		return nil, err
	}

	opts := func(langs []string) []huh.Option[string] {
		var huhOpts []huh.Option[string]
		for _, v := range langs {
			huhOpts = append(huhOpts, huh.Option[string]{
				Key:   v,
				Value: v})
		}
		return huhOpts
	}

	err = huh.
		NewSelect[string]().
		Title("Language to generate `.gitignore` for").
		Options(opts(utils.CommonLangs())...).
		Value(&lang).
		Run()
	if err != nil {
		log.Fatal(err)
	}

	if lang == "Custom" {
		customLang, err := utils.ReadInput("Your Language")
		if err != nil {
			log.Fatal(err)
		}
		lang = cases.Title(language.English, cases.NoLower).String(customLang)
	}

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

	lang, err := utils.ReadInput("Language to setup gitignore for")
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
	commitMsg := "chore(init): create README"
	fContent := "# " + *r.Name
	currentTime := time.Now()

	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return custom.Error(custom.ErrMsg["getGhUserErr"], err)
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
		msg := custom.ErrMsg["fileCreationErr"] + "README.md"
		return custom.Error(msg, err)
	}
	return nil
}

func CreateGitignore(c *github.Client, ctx context.Context, r *github.Repository, lang string) error {
	commitMsg := "chore(init): create .gitignore"
	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return custom.Error(custom.ErrMsg["getGhUserErr"], err)
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
		msg := custom.ErrMsg["fileCreationErr"] + ".gitignore"
		return custom.Error(msg, err)
	}
	return nil
}

func CreateLicense(c *github.Client, ctx context.Context, l string, rName string) error {
	commitMsg := "chore: add LICENSE"
	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return custom.Error(custom.ErrMsg["getGhUserErr"], err)
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
		Content:   []byte(l),
		Author:    &a,
		Committer: &a,
	}
	_, _, err = c.Repositories.CreateFile(ctx, *a.Login, rName, "LICENSE", &fOpts)
	if err != nil {
		msg := custom.ErrMsg["fileCreationErr"] + "LICENSE"
		return custom.Error(msg, err)
	}
	return nil
}

func GetGitignore(c *github.Client, ctx context.Context, lang string) (string, error) {
	gitIgnore, _, err := c.Gitignores.Get(ctx, lang)
	if err != nil {
		return "", custom.Error(custom.ErrMsg["getIgnoreErr"], err)
	}
	ignoreContents := *gitIgnore.Source
	comment := "\n# Incase the template didn't include this for some reason\n"
	if !strings.Contains(ignoreContents, "env") {
		ignoreContents := ignoreContents + comment + ".env"
		return ignoreContents, nil
	}

	return ignoreContents, nil
}
