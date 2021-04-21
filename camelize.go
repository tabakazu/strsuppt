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
