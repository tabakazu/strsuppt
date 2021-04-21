package strsuppt

import "strings"

func Capitalize(s string) string {
	first := string(s[0])
	return strings.Replace(s, first, strings.Title(first), 1)
}
