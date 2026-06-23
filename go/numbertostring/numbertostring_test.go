package main

import (
	"fmt"
	"testing"
)

func TestNumberToString(t *testing.T) {
	var tests = []struct {
		in   int
		want string
	}{
		{67, "67"},
		{79585, "79585"},
		{3, "3"},
		{-1, "-1"},
		{0, "0"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v, %v", test.in, test.want)
		t.Run(name,
			func(t *testing.T) {
				answer := NumberToString(test.in)
				if answer != test.want {
					t.Errorf("got %v, want %v", answer, test.want)
				}
			})
	}
}
