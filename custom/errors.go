package custom

import "errors"

var ErrMsg = map[string]string{
	"noEnvErr":        "Looks like you don't have a .env\nlet's fix that.",
	"emptyEnvErr":     "Your .env is empty\nlet's fix that.",
	"envCreationErr":  "Could not create `.env`.",
	"readInputErr":    "Could not read your input.",
	"emptyInputErr":   "No input was entered.",
	"repoCreationErr": "Could not create repository.",
	"getGhUserErr":    "Could not get the GitHub user.",
	"fileCreationErr": "Could not create file",
	"getIgnoreErr":    "Could not get .gitignore template.",
	"issuesErr":       "Could not get your issues.",
	"dirCreationErr":  "Could not create `.rgn`.",
	"homeErr":         "HOME DIRECTORY SHOULD EXIST.",
	"repoExistsErr":   "Repository already exists.",
}

func Error(msg string, err error) error {
	newErr := errors.New(msg)
	e := errors.Join(newErr, err)
	return e
}
