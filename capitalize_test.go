package strsuppt_test

import (
	"testing"

	"github.com/tabakazu/strsuppt"
)

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
		out := strsuppt.Capitalize(test.in)
		if test.out != out {
			t.Errorf("Capitalize(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}
