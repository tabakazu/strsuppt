package strsuppt_test

import (
	"testing"

	"github.com/tabakazu/strsuppt"
)

func TestCamelize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"string_support", "StringSupport"},
		{"string_support support", "StringSupport Support"},
	}
	for _, test := range tests {
		out := strsuppt.Camelize(test.in)
		if test.out != out {
			t.Errorf("Camelize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}
