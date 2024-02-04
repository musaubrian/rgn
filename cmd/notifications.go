package cmd

import (
	"log"

	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

// notificationsCmd represents the notifications command
var notificationsCmd = &cobra.Command{
	Use:     "notification",
	Short:   "Get a brief overview of your unread notifications",
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := gh.GetUnreadNotifications(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		t := custom.NewTabby()
		t.AddHeader("\n#", "Repo", "Type", "Title")

		for i, v := range r {
			if *v.Unread {
				title := utils.Truncate(*v.Subject.Title, 25)
				t.AddLine(i, *v.Repository.FullName, *v.Subject.Type, title)
			}
		}
		t.Print()

	},
}

var readnotifications = &cobra.Command{
	Use:     "read",
	Short:   "Mark notifications read",
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		err := gh.MarkNotificationsRead(client, cmd.Context())
		if err != nil {
			log.Fatal("Could not mark notifications as read")
		}
	},
}

func init() {
	rootCmd.AddCommand(notificationsCmd)
	notificationsCmd.AddCommand(readnotifications)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notificationsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notificationsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
