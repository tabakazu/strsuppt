package strsuppt

import (
	"regexp"
	"strings"
)

// Camelize converts strings to UpperCamelCase.
func Camelize(s string) string {
	var str string
	ws := strings.Split(s, "_")
	for _, w := range ws {
		str = str + strings.Title(w)
	}

	str = strings.Replace(str, "/", "::", -1)
	return str
}

func Capitalize(s string) string {
	first := string(s[0])
	return strings.Replace(s, first, strings.Title(first), 1)
}

// Underscore convert strings to LowerCase.
func Underscore(s string) string {
	r := regexp.MustCompile("[A-Z]")
	matches := r.FindAllString(s, -1)
	for idx, w := range matches {
		if idx == 0 {
			s = strings.Replace(s, w, strings.ToLower(w), 1)
			continue
		}
		s = strings.Replace(s, w, "_"+strings.ToLower(w), 1)
	}
	s = strings.Replace(s, "::", "/", -1)
	s = strings.Replace(s, "/_", "/", -1)
	return s
}
