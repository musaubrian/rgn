package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/musaubrian/rgn/custom"
)

func ReadEnv(envPath string) ([]string, error) {
	var ghDets []string
	envFile, err := os.Open(envPath)
	if err != nil {
		return ghDets, custom.Error(custom.ErrMsg["noEnvErr"], err)
	}
	defer envFile.Close()

	sc := bufio.NewScanner(envFile)
	for sc.Scan() {
		ghDets = append(ghDets, sc.Text())
	}

	return ghDets, nil
}

func CreateEnv(envPath string) error {
	err := createEnvDir()
	if err != nil {
		return err
	}
	env, err := os.Create(envPath)
	if err != nil {
		return custom.Error(custom.ErrMsg["envCreationErr"], err)
	}
	defer env.Close()
	fmt.Println("// Let's set you up")
	username, err := ReadInput("Your Github username:")
	if err != nil {
		return err
	}
	token, err := ReadInput("Your Token(PAT):")
	if err != nil {
		return err
	}

	env.WriteString(fmt.Sprintf("username=%s\ntoken=%s\n", username, token))

	fmt.Println("// All set, re-run to access the commands")

	return nil
}

func UpdateEnv(envPath string) error {
	dets, err := ReadEnv(envPath)
	if err != nil {
		return err
	}

	username := strings.Split(dets[0], "=")[1]
	token, err := ReadInput("New token:")
	if token == "" {
		log.Fatal("Empty token, Ignoring")
	}
	if err != nil {
		return err
	}
	f, err := os.OpenFile(envPath, os.O_WRONLY|os.O_TRUNC, 0660)
	defer f.Close()
	if err != nil {
		return err
	}
	p := fmt.Sprintf("username=%s\ntoken=%s\n", username, token)
	f.WriteString(p)

	return nil
}

func createEnvDir() error {
	h, err := os.UserHomeDir()
	if err != nil {
		return custom.Error(custom.ErrMsg["homeErr"], nil)
	}
	envDir := h + string(os.PathSeparator) + ".rgn"
	err = os.Mkdir(envDir, 0o770)
	if err != nil {
		return custom.Error(custom.ErrMsg["dirCreationErr"], err)
	}
	return nil
}

func GetEnvLoc() (string, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return "", custom.Error(custom.ErrMsg["homeErr"], nil)
	}
	envLoc := h + string(os.PathSeparator) + ".rgn" + string(os.PathSeparator) + ".env"
	return envLoc, nil
}
