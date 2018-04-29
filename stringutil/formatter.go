package stringutil

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase converts camel case to snake case
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// Decapitalize decapitalize input string
func Decapitalize(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

// Capitalize capitalize input string
func Capitalize(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}
