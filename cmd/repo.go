/*
Copyright © 2023 musaubrian
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Create a new empty repo",
	Long: `Create a new empty repo (no README or .gitignore)
It displays the URLs you can use to clone it(SSH and HTTPS) when done

You can also configure it to some level:
- Completely bare-bones(no README or .gitignore)
    run: rgn repo --help for more information
- Just a README.
    run: rgn repo r --help for more information
- Just a .gitignore of the specified language
    run: rgn repo i --help for more information
- Full.
    run: rgn repo a --help for more information
`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := gh.CreateEmptyRepo(Client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("// Created %s successfully\n", *r.Name)
		fmt.Printf("Cloning URLs\nssh: %s\nhttps: %s", *r.SSHURL, *r.CloneURL)
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
		r, err := gh.CreateRepoWithGitignore(Client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("// Created %s successfully\n", *r.Name)
		fmt.Printf("Cloning URLs\nssh: %s\nhttps: %s", *r.SSHURL, *r.CloneURL)
	},
}

var withReadme = &cobra.Command{
	Use:   "readme",
	Short: "Create a github repo with only a Readme.md",
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
		r, err := gh.CreateRepoWithReadme(Client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("// Created %s successfully\n", *r.Name)
		fmt.Printf("Cloning URLs\nssh: %s\nhttps: %s", *r.SSHURL, *r.CloneURL)
	},
}

var readmeAndGitignore = &cobra.Command{
	Use:     "all",
	Short:   "Create a github repo with both a README and .gitignore",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := gh.CreateRepoWithBoth(Client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("// Created %s successfully\n", *r.Name)
		fmt.Printf("Cloning URLs\nssh: %s\nhttps: %s", *r.SSHURL, *r.CloneURL)
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)
	repoCmd.AddCommand(withGitignore)
	repoCmd.AddCommand(withReadme)
	repoCmd.AddCommand(readmeAndGitignore)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
