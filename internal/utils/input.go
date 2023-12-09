package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/musaubrian/rgn/custom"
)

func ReadInput(prompt string) (string, error) {
	var s string
	var err error

	rd := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)
	s, err = rd.ReadString('\n')

	if err != nil {
		return s, custom.Error(custom.ErrMsg["readInputErr"], err)
	}

	s = strings.TrimSuffix(s, "\n")
	if len(s) < 1 {
		return s, custom.Error(custom.ErrMsg["emptyInputErr"], nil)
	}

	return s, nil
}
