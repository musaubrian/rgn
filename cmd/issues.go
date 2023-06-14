package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List issues related to the user",
	/* Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`, */

	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RUN: rgn issues --help for more information")
	},
}

var assignIssues = &cobra.Command{
	Use:   "assigned",
	Short: "List all open issues assigned to you",
	Long: `List all open issues assigned to you
Response includes any labels available`,
	Run: func(cmd *cobra.Command, args []string) {
		t := tabby.New()
		t.AddHeader("\nNO.", "REPO", "TITLE", "BODY", "LABELS", "CREATED_AT")
		issues, err := gh.GetIssuesAssigned(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		for _, i := range issues {
			b := strings.Split(*i.Body, "\r\n")
			body := b[0]

			// only when there was more than one line
			if len(b) > 1 {
				body = body + "..."
			}
			timePassed := time.Since(*i.CreatedAt)
			relativeTime := timePassed.String() + " ago"

			if len(i.Labels) > 0 {
				l := *i.Labels[0].Name

				if len(i.Labels) > 1 {
					l = l + "..."
				}

				t.AddLine(count, *i.Repository.FullName, *i.Title, body, l, relativeTime)
				count += 1
			} else {
				t.AddLine(count, *i.Repository.FullName, *i.Title, body, "--none--", relativeTime)
				count += 1
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
