package strsuppt

import "strings"

func Camelize(s string) string {
	var str string
	ws := strings.Split(s, "_")
	for _, w := range ws {
		str = str + strings.Title(w)
	}
	return str
}

func Capitalize(s string) string {
	first := string(s[0])
	return strings.Replace(s, first, strings.Title(first), 1)
}
