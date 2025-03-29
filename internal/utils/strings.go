package utils

import "strings"

func HeadToLower(str string) string {
	if len(str) <= 1 {
		return strings.ToLower(str)
	}

	return strings.ToLower(str[:1]) + str[1:]
}

func HeadToUpper(str string) string {
	if len(str) <= 1 {
		return strings.ToUpper(str)
	}

	return strings.ToUpper(str[:1]) + str[1:]
}
