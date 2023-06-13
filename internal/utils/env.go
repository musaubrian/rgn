package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadEnv(envPath string) ([]string, error) {
	var ghDets []string
	envFile, err := os.Open(envPath)
	if err != nil {
		return ghDets, err
	}
	defer envFile.Close()

	sc := bufio.NewScanner(envFile)
	for sc.Scan() {
		ghDets = append(ghDets, sc.Text())
	}

	return ghDets, nil
}

func CreateEnv(envPath string) error {
	var envDets []string

	env, err := os.Create(envPath)
	if err != nil {
		return err
	}
	defer env.Close()
	fmt.Println("// Let's set you up")
	uName, err := ReadInput("Your Github username:")
	if err != nil {
		return err
	}
	uName = "username=" + uName
	envDets = append(envDets, uName)
	token, err := ReadInput("Your Token(PAT):")
	if err != nil {
		return err
	}
	token = "token=" + token

	envDets = append(envDets, token)

	for _, v := range envDets {
		env.WriteString(v + "\n")
	}
	fmt.Println("// All set, re-run to access the commands")

	return nil
}
