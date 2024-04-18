package cmd

import (
	"log"
	"strings"

	"github.com/charmbracelet/huh/spinner"
	"github.com/musaubrian/rgn/custom"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update credentials",
	Long: `
Update your token when it expires`,
	Aliases: []string{"u"},
	Run: func(cmd *cobra.Command, args []string) {
		env, err := utils.GetEnvLoc()
		if err != nil {
			log.Fatal(err)
		}
		err = utils.UpdateToken(env)
		if err != nil {
			log.Fatal(err)
		}
		ping := func() {
			if err := gh.Ping(client, rootCmd.Context()); err != nil {
				if strings.Contains(err.Error(), "401") {
					custom.HeaderMsg("Invalid token")
					log.Fatal(err)
				}
			}
		}
		err = spinner.New().Title("Checking token validity...").Action(ping).Run()
		if err != nil {
			log.Fatal(err)
		}
		custom.Success("Token updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
