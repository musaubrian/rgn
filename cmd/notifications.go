package cmd

import (
	"log"

	"github.com/cheynewallace/tabby"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

// notificationsCmd represents the notifications command
var notificationsCmd = &cobra.Command{
	Use:   "notification",
	Short: "Get a brief overview of your unread notifications",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:
	//
	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := gh.GetUnreadNotifications(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		t := tabby.New()
		t.AddHeader("\n#", "Repo", "Type", "Title")

		for i, v := range r {
			if *v.Unread {
				title := utils.Truncate(*v.Subject.Title, 50)
				t.AddLine(i, *v.Repository.FullName, *v.Subject.Type, title)
			}
		}
		t.Print()

	},
}

func init() {
	rootCmd.AddCommand(notificationsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notificationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notificationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
