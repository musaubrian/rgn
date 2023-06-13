package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func ReadInput(prompt string) (string, error) {
	var s string
	var err error

	rd := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", prompt)
	s, err = rd.ReadString('\n')

	if err != nil {
		return s, err
	}

	s = strings.TrimSuffix(s, "\n")
	if len(s) < 1 {
		return s, errors.New("No input entered")
	}

	return s, nil
}
