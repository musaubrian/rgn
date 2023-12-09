package utils

import (
	"fmt"
	"time"
)

func TimeInDays(t time.Duration) string {
	var relTime string
	h := int(t.Hours())
	m := int(t.Minutes())
	d := h / 24
	hrs := h % 24

	if hrs == 0 && m > 0 && m < 1440 {
		relTime = fmt.Sprintf("%d mins", m)
		return relTime
	}
	relTime = fmt.Sprintf("%dd %dhrs", d, hrs)
	return relTime
}
