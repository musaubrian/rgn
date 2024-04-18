package utils

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func notEmpty(s string) error {
	if s == "" {
		return errors.New("Empty values are not permmitted")
	}
	return nil
}

func ReadInput(prompt string) (string, error) {
	var s string
	err := huh.NewInput().
		Title(prompt).
		Value(&s).
		Validate(notEmpty).
		Run()

	if err != nil {
		return s, err
	}

	return s, nil
}
