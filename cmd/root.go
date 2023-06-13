/*
Copyright © 2023 musaubrian
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
)

var Client *github.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rgn",
	Short: "A lightweight version of Github CLI",
	Long: `
repoGen(rgn) gives you the ability to create repositories on your github account from the command line.

You can configure the way you want the repo to be.
run: rgn repo --help for more information
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

	Client, err = gh.Auth(".env", rootCmd.Context())
	if err != nil {
		log.Fatal(err)
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
