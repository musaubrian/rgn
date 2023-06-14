package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RUN: rgn issues --help for more information")
	},
}

var assignIssues = &cobra.Command{
	Use:  "assigned",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		t := tabby.New()
		t.AddHeader("\nNO.", "REPO", "TITLE", "BODY", "LABELS", "ASSIGNEES", "LOCKED", "CREATED_AT")
		issues, err := gh.GetIssuesAssigned(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		for _, i := range issues {
			for _, l := range i.Labels {
				for _, a := range i.Assignees {
					timePassed := time.Since(*i.CreatedAt)
					relativeTime := timePassed.String() + " ago"
					t.AddLine(count, *i.Repository.FullName, *i.Title, *i.Body, *l.Name, *a.Login, *i.Locked, relativeTime)
					count += 1
				}
			}
		}
		t.Print()
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)
	issuesCmd.AddCommand(assignIssues)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
