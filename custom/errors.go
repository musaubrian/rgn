package custom

import "errors"

func NoneExistentEnvErr() error {
	return errors.New("Looks like you don't have a .env\nLet's fix that")
}

func EmptyEnvErr() error {
	return errors.New("Your .env is empty\nLet's fix that")
}

func EnvCreationErr(err error) error {
	envCreationErr := errors.New("Could not create .env\n")
	e := errors.Join(envCreationErr, err)
	return e
}

func ReadInputErr() error {
	return errors.New("Could not read your input")
}

func EmptyInputReadErr() error {
	return errors.New("No input was entered")
}

func GetGHUserErr(err error) error {
	ghUserErr := errors.New("Could not get the github user\n")
	e := errors.Join(ghUserErr, err)
	return e
}

func CreateRepoErr(err error) error {
	createRepoErr := errors.New("Could not create repository.\n")

	e := errors.Join(createRepoErr, err)
	return e
}

func FileCreationErr(f string, err error) error {
	msg := "Could not create " + f + "\n"
	fileCreationErr := errors.New(msg)
	e := errors.Join(fileCreationErr, err)
	return e
}

func GetGitignoreErr(err error) error {
	getIgnoreErr := errors.New("Could not get .gitignore template\n")
	e := errors.Join(getIgnoreErr, err)
	return e
}
