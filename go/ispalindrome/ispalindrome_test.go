package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		{"a", true},
		{"aba", true},
		{"Abba", true},
		{"hello", false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%s, %t", test.in, test.want)
		t.Run(name,
			func(t *testing.T) {
				answer := IsPalindrome(test.in)
				if answer != test.want {
					t.Errorf("got %t, want %t", answer, test.want)
				}
			})
	}
}
