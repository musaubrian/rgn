package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/musaubrian/rgn/custom"
)

func ReadEnv(envPath string) ([]string, error) {
	var ghDets []string
	envFile, err := os.Open(envPath)
	if err != nil {
		return ghDets, custom.NoneExistentEnvErr()
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

	err := createEnvDir()
	if err != nil {
		return err
	}
	env, err := os.Create(envPath)
	if err != nil {
		return custom.EnvCreationErr(err)
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

func createEnvDir() error {
	h, err := os.UserHomeDir()
	if err != nil {
		return custom.GetHomeErr(err)
	}
	envDir := h + string(os.PathSeparator) + ".rgn"
	err = os.Mkdir(envDir, 0o770)
	if err != nil {
		return custom.DirCreationErr(err)
	}
	return nil
}

func GetEnvLoc() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", custom.GetHomeErr(err)
	}
	envLoc := h + string(os.PathSeparator) + ".rgn" + string(os.PathSeparator) + ".env"
	return envLoc, nil
}
