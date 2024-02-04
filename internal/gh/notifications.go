package gh

import (
	"context"
	"time"

	"github.com/google/go-github/github"
)

func GetUnreadNotifications(c *github.Client, ctx context.Context) ([]*github.Notification, error) {
	opts := github.NotificationListOptions{
		All:           true,
		Participating: false,
		ListOptions: github.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	res, _, err := c.Activity.ListNotifications(ctx, &opts)
	return res, err
}

func MarkNotificationsRead(c *github.Client, ctx context.Context) error {
	_, err := c.Activity.MarkNotificationsRead(ctx, time.Now())
	return err
}
