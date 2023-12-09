package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List all open issues assigned to you",
	Long: `List all open issues assigned to you
Response includes any labels available`,
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		var b []string
		var body string
		t := custom.NewCustomTabby()
		t.AddHeader("\nNO.", "REPO", "TITLE", "BODY", "LABELS", "CREATED_AT")
		issues, err := gh.GetIssuesAssigned(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		for c, i := range issues {
			if len(i.GetBody()) > 1 {
				if strings.Contains(*i.Body, "\r\n") {
					b = strings.Split(*i.Body, "\r\n")
				}

				b = strings.Split(*i.Body, ".")
				body = b[0]
			}
			body = "--Nothin here--"
			// only when there was more than one line
			if len(b) > 1 {
				body = utils.Truncate(body, len(body)-len(body)/3)
			}

			timePassed := time.Since(*i.CreatedAt)
			relativeTime := utils.TimeInDays(timePassed) + " ago"

			if len(i.Labels) > 0 {
				l := *i.Labels[0].Name

				if len(i.Labels) > 1 {
					l = l + "..."
				}

				t.AddLine(c, *i.Repository.FullName, *i.Title, body, l, relativeTime)
				c += 1
			} else {
				t.AddLine(c, *i.Repository.FullName, *i.Title, body, "--none--", relativeTime)
				c += 1
			}
		}
		t.Print()
	},
}

func init() {
	rootCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
