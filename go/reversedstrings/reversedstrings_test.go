package reversedstrings

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"world", "dlrow"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s, %s", test.in, test.want)
		t.Run(name,
			func(t *testing.T) {
				answer := Reverse(test.in)
				if answer != test.want {
					t.Errorf("got %s, want %s", answer, test.want)
				}
			})
	}
}
