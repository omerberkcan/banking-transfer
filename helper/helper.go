package helper

import (
	"regexp"
)

func IsNumeric(s string) bool {
	// Define a regular expression pattern to match numeric characters only
	pattern := "^[0-9]+$"

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Use MatchString to check if the string matches the pattern
	return regex.MatchString(s)
}
