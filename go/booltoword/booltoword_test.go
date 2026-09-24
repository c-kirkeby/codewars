package booltoword

import (
	"fmt"
	"testing"
)

func TestBoolToWord(t *testing.T) {
	var tests = []struct {
		in   bool
		want string
	}{
		{true, "Yes"},
		{false, "No"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%t, %s", test.in, test.want)
		t.Run(name,
			func(t *testing.T) {
				answer := BoolToWord(test.in)
				if answer != test.want {
					t.Errorf("got %s, want %s", answer, test.want)
				}
			})
	}
}
