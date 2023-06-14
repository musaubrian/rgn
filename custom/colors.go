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

func SuccesfullCreation(rName, sshURL, httpURL string) {
	success.Printf("// Created [%s] successfully\n", rName)
	bold.Println("\n// Cloning URLs")
	fmt.Printf("ssh: %s\nhttps: %s", sshURL, httpURL)
}
