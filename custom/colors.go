package custom

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	bold    = color.New(color.Bold)
	success = color.New(color.Bold, color.FgGreen)
)

func HeaderMsg(msg string) {
	bold.Println("\n// ", msg)
}
func Success(msg string) {
	success.Printf("%s", msg)
}

func SuccesfullRepoCreation(rName, sshURL, httpURL string) {
	Success(fmt.Sprintf("// Created [%s] successfully\n", rName))
	bold.Println("\n// Cloning URLs")
	fmt.Printf("ssh: %s\nhttps: %s", sshURL, httpURL)
}

func SuccesfullLicenseCreation(rName string) {
	Success(fmt.Sprintf("// Created LICENSE in [%s]\n", rName))
}
