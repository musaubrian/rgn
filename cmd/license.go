/*
Copyright © 2023 musaubrian
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/musaubrian/rgn/internal/gh"
	"github.com/musaubrian/rgn/internal/utils"
	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Add a LICENSE to your project",
	Long: `Generates a LICENSE for use in your project

LICENSE can either be added to the repo that is already on github
or the contents of the LICENSE can be printed out in the stdout locally

To learn more about which LICENSE to choose, visit 
https://choosealicense.com/licenses/
    `,
	Aliases: []string{"l"},
}

var mitLicense = &cobra.Command{
	Use:   "mit",
	Short: "Add an MIT LICENSE to your project",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}
		l, err := gh.GetMIT(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		str := *l.Body
		cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
		if err != nil {
			log.Fatal(err)
		}
		if g {
            rName, err := utils.ReadInput("Repository to create LICENSE for:")
			if err != nil {
				log.Fatal(err)
			}
			gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
		} else {
			fmt.Println(cleanStr)
		}
	},
}

var apacheLicense = &cobra.Command{
	Use:   "apache",
	Short: "Add an Apache License 2.0",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}
		l, err := gh.GetApache(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		str := *l.Body
		cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
		if err != nil {
			log.Fatal(err)
		}
		if g {
			rName, err := utils.ReadInput("Repository to create LICENSE for:")
			if err != nil {
				log.Fatal(err)
			}
			gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
		} else {
			fmt.Println(cleanStr)
		}
	},
}

var bsdLicense = &cobra.Command{
	Use:   "bsd",
	Short: "Add BSD License (v2 or v3)",
	Long: `Create a new BSD License
rgn license bsd 2/3
- 2 to use BSD 2-Clause
- 3 to use BSD 3-Clause
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
            log.Fatal("You need to pick a version\n\nRun: rgn l bsd -h for more info")
		}
		v, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}

		if v == 2 {
			l, err := gh.GetBSD2(client, cmd.Context())
			if err != nil {
				log.Fatal(err)
			}
			str := *l.Body
			cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
			if err != nil {
				log.Fatal(err)
			}
			if g {
				rName, err := utils.ReadInput("Repository to create LICENSE for:")
				if err != nil {
					log.Fatal(err)
				}
				gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
			} else {
				fmt.Println(cleanStr)
			}
		} else if v == 3 {

			l, err := gh.GetBSD3(client, cmd.Context())
			if err != nil {
				log.Fatal(err)
			}
			str := *l.Body
			cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
			if err != nil {
				log.Fatal(err)
			}
			if g {
				rName, err := utils.ReadInput("Repository to create LICENSE for:")
				if err != nil {
					log.Fatal(err)
				}
				gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
			} else {
				fmt.Println(cleanStr)
			}
		}
    },
}

var gplLicense = &cobra.Command{
	Use:   "gpl",
	Short: "Add GNU General Public License (v2.0 or v3.0)",
	Long: `Create a new BSD license

rgn license bsd 2/3

- 2 to use GNU General Public License v2.0
- 3 to use GNU General Public License v3.0
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
            log.Fatal("You need to pick a version\n\nRun: rgn l gpl -h for more info")
		}
		v, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}

		if v == 2 {
			l, err := gh.GetGPL2(client, cmd.Context())
			if err != nil {
				log.Fatal(err)
			}
			str := *l.Body
			cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
			if err != nil {
				log.Fatal(err)
			}
			if g {
				rName, err := utils.ReadInput("Repository to create LICENSE for:")
				if err != nil {
					log.Fatal(err)
				}
				gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
			} else {
				fmt.Println(cleanStr)
			}
		} else if v == 3 {

			l, err := gh.GetGPL3(client, cmd.Context())
			if err != nil {
				log.Fatal(err)
			}
			str := *l.Body
			cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
			if err != nil {
				log.Fatal(err)
			}
			if g {
				rName, err := utils.ReadInput("Repository to create LICENSE for:")
				if err != nil {
					log.Fatal(err)
				}
				gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
			} else {
				fmt.Println(cleanStr)
			}
		}
	},
}

var cc0License = &cobra.Command{
	Use:   "cc",
	Short: "Add Creative COmmons Zero v1.0 Universal",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}
		l, err := gh.GetCC0(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		str := *l.Body
		cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
		if err != nil {
			log.Fatal(err)
		}
		if g {
			rName, err := utils.ReadInput("Repository to create LICENSE for:")
			if err != nil {
				log.Fatal(err)
			}
			gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
		} else {
			fmt.Println(cleanStr)
		}
	},
}

var mplLicense = &cobra.Command{
	Use:   "mpl",
	Short: "Add Mozilla Public License v2.0",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}
		l, err := gh.GetMPL(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		str := *l.Body
		cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
		if err != nil {
			log.Fatal(err)
		}
		if g {
			rName, err := utils.ReadInput("Repository to create LICENSE for:")
			if err != nil {
				log.Fatal(err)
			}
			gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
		} else {
			fmt.Println(cleanStr)
		}
	},
}

var lgplLicense = &cobra.Command{
	Use:   "lgpl",
	Short: "Add GNU Lesser General Public License v2.1",
	Run: func(cmd *cobra.Command, args []string) {
		g, err := licenseCmd.Flags().GetBool("github")
		if err != nil {
			log.Fatal(err)
		}
		l, err := gh.GetLGPL(client, cmd.Context())
		if err != nil {
			log.Fatal(err)
		}
		str := *l.Body
		cleanStr, err := utils.FillLicense(client, cmd.Context(), str)
		if err != nil {
			log.Fatal(err)
		}
		if g {
			rName, err := utils.ReadInput("Repository to create LICENSE for:")
			if err != nil {
				log.Fatal(err)
			}
			gh.CreateLicense(client, cmd.Context(), cleanStr, rName)
		} else {
			fmt.Println(cleanStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)
	licenseCmd.AddCommand(mitLicense)
	licenseCmd.AddCommand(apacheLicense)
	licenseCmd.AddCommand(lgplLicense)
	licenseCmd.AddCommand(mplLicense)
	licenseCmd.AddCommand(bsdLicense)
	licenseCmd.AddCommand(cc0License)
	licenseCmd.AddCommand(gplLicense)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// licenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	licenseCmd.PersistentFlags().BoolP("github", "g", false, "Create the LICENSE on github instead of just printing it out to stdout")
}
