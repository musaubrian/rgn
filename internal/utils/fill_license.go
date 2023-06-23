package utils

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/google/go-github/github"
	"github.com/musaubrian/rgn/custom"
)

// Replace [year] and [fullname] in the License with actual values
func FillLicense(c *github.Client, ctx context.Context, l string) (string, error) {
	var cleanLicense string
	u, _, err := c.Users.Get(ctx, "")
	if err != nil {
		return cleanLicense, custom.GetGHUserErr(err)
	}

	// convert yaer to string
	year := fmt.Sprintf("%d", time.Now().Year())

	userPattern := regexp.MustCompile("\\[fullname]")
	altUserPattern := regexp.MustCompile("\\[name of copyright owner]")
	yearPattern := regexp.MustCompile("\\[year]")
	altyearPattern := regexp.MustCompile("\\[yyyy]")

	cleanLicense = userPattern.ReplaceAllString(l, *u.Login)
	cleanLicense = yearPattern.ReplaceAllString(cleanLicense, year)
	cleanLicense = altUserPattern.ReplaceAllString(l, *u.Login)
	cleanLicense = altyearPattern.ReplaceAllString(cleanLicense, year)

	return cleanLicense, nil
}
