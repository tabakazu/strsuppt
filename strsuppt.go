package strsuppt

import (
	"regexp"
	"strconv"
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

// Ordinal returns suffix of number to denote the position in an ordered
func Ordinal(s string) string {
	if _, err := strconv.Atoi(s); err != nil {
		return ""
	}
	switch s {
	case "1":
		return "st"
	case "2":
		return "nd"
	case "3":
		return "rd"
	case "4", "5", "6", "7", "8", "9", "10", "11", "12", "13":
		return "th"
	default:
		switch s[len(s)-1:] {
		case "1":
			return "st"
		case "2":
			return "nd"
		case "3":
			return "rd"
		default:
			return "th"
		}
	}
}

func Ordinalize(s string) string {
	return s + Ordinal(s)
}

// Dasherize returns replace underscore to dash in string
func Dasherize(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

func Deconstantize(s string) string {
	idx := strings.LastIndex(s, "::")
	if idx < 0 {
		return ""
	}
	return s[0:idx]
}
