package utils

import "unicode"

func Truncate(text string, maxLen int) string {
	lastSpaceIdx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIdx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIdx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}
