package utils_test

import (
	"github.com/suger-131997/dein/internal/utils"
	"testing"
)

func TestHeadToLower(t *testing.T) {
	tests := []struct {
		name string

		in string

		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "single letter",
			in:   "A",
			want: "a",
		},
		{
			name: "multiple letters",
			in:   "ABC",
			want: "aBC",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := utils.HeadToLower(tc.in)
			if got != tc.want {
				tt.Errorf("HeadToLower(%q) = %q; want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestHeadToUpper(t *testing.T) {
	tests := []struct {
		name string

		in string

		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "",
		},
		{
			name: "single letter",
			in:   "a",
			want: "A",
		},
		{
			name: "multiple letters",
			in:   "abc",
			want: "Abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			got := utils.HeadToUpper(tc.in)
			if got != tc.want {
				tt.Errorf("HeadToUpper(%q) = %q; want %q", tc.in, got, tc.want)
			}
		})
	}
}
