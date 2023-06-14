package gh

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/internal/utils"
	"golang.org/x/oauth2"
)

func Auth(envPath string, c context.Context) (*github.Client, error) {
	var client *github.Client
	var tkn []string

	envDets, err := utils.ReadEnv(envPath)
	if err != nil {
		err := utils.CreateEnv(envPath)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	if len(envDets) != 2 {
		log.Println("It seems your .env is empty\nLet's fix that")
		err := utils.CreateEnv(envPath)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)

	}
	tkn = strings.Split(envDets[1], "=")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tkn[1]},
	)
	tc := oauth2.NewClient(c, ts)
	client = github.NewClient(tc)
	return client, nil
}
