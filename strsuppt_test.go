package strsuppt

import (
	"testing"
)

func TestCamelize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"string_support", "StringSupport"},
		{"string_support/support", "StringSupport::Support"},
	}
	for _, test := range tests {
		out := Camelize(test.in)
		if test.out != out {
			t.Errorf("Camelize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"taba", "Taba"},
		{"taba san", "Taba san"},
		{"taba san san", "Taba san san"},
		{"0san", "0san"},
	}
	for _, test := range tests {
		out := Capitalize(test.in)
		if test.out != out {
			t.Errorf("Capitalize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}
