package strsuppt

import "testing"

func TestCamelize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "string_support"}, "StringSupport"},
		{"", args{s: "string_support support"}, "StringSupport Support"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camelize(tt.args.s); got != tt.want {
				t.Errorf("Camelize() = %v, want %v", got, tt.want)
			}
		})
	}
}
