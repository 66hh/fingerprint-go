package utils

import "strings"

func FindText(text string, substr string) bool {
	if strings.Index(strings.ToLower(text), strings.ToLower(substr)) != -1 {
		return true
	}
	return false
}
