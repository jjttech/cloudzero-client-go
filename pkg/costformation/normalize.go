package costformation

import (
	"regexp"
	"strings"
)

// Normalize This transform will lowercase all of the letters in a the source value and will also convert any instance of the characters .,/#!$%^&*;:=_~()\' to a dash (-).
func Normalize(in string) string {
	re := regexp.MustCompile(`[.,/#!$%^&\*;:=_+~()\\'\s]`)

	// Remove leading/trailing spaces
	// lower case
	// replace special chars via regexp
	return re.ReplaceAllString(strings.ToLower(strings.Trim(in, " \t")), "-")
}
