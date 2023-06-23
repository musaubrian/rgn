package gh

import (
	"context"

	"github.com/google/go-github/github"
)

func GetApache(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "apache-2.0")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetBSD2(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "bsd-2-clause")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetBSD3(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "bsd-3-clause")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetMIT(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "mit")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetBSL(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "bsl-1.0")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetCC0(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "cc0-1.0")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetGPL2(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "gpl-2.0")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetGPL3(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "gpl-30")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetLGPL(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "lgpl-2.1")
	if err != nil {
		return l, err
	}
	return l, err
}

func GetMPL(c *github.Client, ctx context.Context) (*github.License, error) {
	l, _, err := c.Licenses.Get(ctx, "mpl-2.0")
	if err != nil {
		return l, err
	}
	return l, err
}
