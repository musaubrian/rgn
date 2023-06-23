/*
Copyright © 2023 musaubrian
*/
package cmd

import (
	"log"

	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new repo",
	Long: `
Create a new repository with the specified configurations
It displays the URLs you can use to clone it(SSH and HTTPS) when done
`,
	Aliases: []string{"c"},
}

var emptyRepo = &cobra.Command{
	Use:   "empty",
	Short: "Create a new empty repo",
	Long: `
Creates a new repository that is completely empty
No README or .gitignore`,
	Aliases: []string{"e"},
	Run: func(cmd *cobra.Command, args []string) {
		custom.HeaderMsg("Creating an empty repo")
		r, err := gh.CreateEmptyRepo(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}

		custom.SuccesfullRepoCreation(*r.Name, *r.SSHURL, *r.CloneURL)
	},
}

var withGitignore = &cobra.Command{
	Use:   "ignore",
	Short: "Create a new repo with only a .gitignore",
	Long: `
Create a new repository with only a .gitignore of the specified language
For full list of languages supported, visit: https://github.com/github/gitignore
It displays the URLs you can use to clone it(SSH and HTTPS)
`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		custom.HeaderMsg("Creating repo with gitignore")
		r, err := gh.CreateRepoWithGitignore(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		custom.SuccesfullRepoCreation(*r.Name, *r.SSHURL, *r.CloneURL)
	},
}

var withReadme = &cobra.Command{
	Use:   "readme",
	Short: "Create a github repo with only a README",
	Long: `
Create a repository with only a README.md
The README.md contains the repository name and the description

-------README.md---------
# Repository name
> Repository description
-------------------------

`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		custom.HeaderMsg("Creating repo with README")
		r, err := gh.CreateRepoWithReadme(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		custom.SuccesfullRepoCreation(*r.Name, *r.SSHURL, *r.CloneURL)
	},
}

var readmeAndGitignore = &cobra.Command{
	Use:     "all",
	Short:   "Create a github repo with both a README and .gitignore",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		custom.HeaderMsg("Creating repo with README & .gitignore")
		r, err := gh.CreateRepoWithBoth(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		custom.SuccesfullRepoCreation(*r.Name, *r.SSHURL, *r.CloneURL)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(emptyRepo)
	createCmd.AddCommand(withGitignore)
	createCmd.AddCommand(withReadme)
	createCmd.AddCommand(readmeAndGitignore)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
