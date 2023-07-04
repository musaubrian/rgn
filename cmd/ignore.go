package cmd

import (
	"fmt"
	"log"

	"github.com/musaubrian/rgn/internal/gh"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ignoreCmd represents the ignore command
var ignoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Print the gitignore contents to stdout",
	Long: `Print the contents of gitignore to stdout
You can redirect the contents to .gitignore
`,
	Example: `rgn ignore go
rgn i go

rgn ignore go > .gitignore
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("You did not include a language\nRun rgn ignore -h")
		} else if len(args) >= 2 {
			log.Fatal("Too many arguments\nRun rgn ignore -h")
		}
		lang := args[0]
		lang = cases.Title(language.English, cases.NoLower).String(lang)
		ignoreContents, err := gh.QuickIgnore(client, cmd.Context(), lang)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ignoreContents)
	},
}

func init() {
	rootCmd.AddCommand(ignoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ignoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ignoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
