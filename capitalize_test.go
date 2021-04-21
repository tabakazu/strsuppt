package strsuppt

import "testing"

func TestCapitalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "taba"}, "Taba"},
		{"", args{s: "taba san"}, "Taba san"},
		{"", args{s: "taba san san"}, "Taba san san"},
		{"", args{s: "0san"}, "0san"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args.s); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
