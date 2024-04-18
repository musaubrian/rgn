package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
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
	Aliases: []string{"ig"},
	Run: func(cmd *cobra.Command, args []string) {
		var lang string

		opts := func(langs []string) []huh.Option[string] {
			var huhOpts []huh.Option[string]
			for _, v := range langs {
				huhOpts = append(huhOpts, huh.Option[string]{
					Key:   v,
					Value: v})
			}
			return huhOpts
		}

		err := huh.
			NewSelect[string]().
			Title("Language to generate `.gitignore` for").
			Options(opts(utils.CommonLangs())...).
			Value(&lang).
			Run()
		if err != nil {
			log.Fatal(err)
		}

		if lang == "Custom" {
			customLang, err := utils.ReadInput("Your Language")
			if err != nil {
				log.Fatal(err)
			}
			lang = cases.Title(language.English, cases.NoLower).String(customLang)
		}

		ignore := func() {
			ignoreContents, err := gh.QuickIgnore(client, cmd.Context(), lang)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ignoreContents)
		}
		err = spinner.New().Title(fmt.Sprintf("Getting %s's `.gitignore`", lang)).Action(ignore).Run()
		if err != nil {
			log.Fatal(err)
		}

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
