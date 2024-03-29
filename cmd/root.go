package cmd

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

var client *github.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rgn",
	Short: "A lightweight version of Github CLI",
	Long: `repoGen(rgn) gives you the ability to create repositories
on your github account from the command line.
It also provides a central point to see issues assigned to you

    `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var err error
	ctx := context.Background()
	rootCmd.SetContext(ctx)
	env, err := utils.GetEnvLoc()
	if err != nil {
		log.Fatal(err)
	}

	client, err = gh.Auth(env, rootCmd.Context())
	if err != nil {
		log.Fatal(err)
	}
	if err := gh.Ping(client, rootCmd.Context()); err != nil {
		if strings.Contains(err.Error(), "401") {
			custom.HeaderMsg("Invalid token")
			err := utils.UpdateToken(env)
			if err != nil {
				log.Println(err)
			}
			// Re-run auth after updating token
			client, err = gh.Auth(env, rootCmd.Context())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rgn.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
