package strsuppt_test

import (
	"fmt"

	"github.com/tabakazu/strsuppt"
)

func ExampleCamelize() {
	s := strsuppt.Camelize("string_support")
	fmt.Println(s)
}

func ExampleCapitalize() {
	s := strsuppt.Capitalize("strsuppt")
	fmt.Println(s)
}
