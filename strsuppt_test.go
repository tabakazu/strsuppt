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

func TestUnderscore(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"StringSupport", "string_support"},
		{"StringSupport::Support", "string_support/support"},
	}
	for _, test := range tests {
		out := Underscore(test.in)
		if test.out != out {
			t.Errorf("Underscore(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}

func TestOrdinal(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "st"},
		{"-1", "st"},
		{"21", "st"},
		{"2", "nd"},
		{"102", "nd"},
		{"3", "rd"},
		{"443", "rd"},
		{"4", "th"},
		{"11", "th"},
		{"-444", "th"},
		{"first", ""},
		{"a1", ""},
	}
	for _, test := range tests {
		out := Ordinal(test.in)
		if test.out != out {
			t.Errorf("Ordinal(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}

func TestOrdinalize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "1st"},
		{"-1", "-1st"},
		{"21", "21st"},
		{"2", "2nd"},
		{"102", "102nd"},
		{"3", "3rd"},
		{"443", "443rd"},
		{"4", "4th"},
		{"11", "11th"},
		{"-444", "-444th"},
		{"first", "first"},
		{"a1", "a1"},
	}
	for _, test := range tests {
		out := Ordinalize(test.in)
		if test.out != out {
			t.Errorf("Ordinalize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}

func TestDasherize(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"foo_bar", "foo-bar"},
		{"__", "--"},
	}
	for _, test := range tests {
		out := Dasherize(test.in)
		if test.out != out {
			t.Errorf("Dasherize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}
