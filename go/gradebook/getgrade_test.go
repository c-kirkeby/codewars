package main

import (
	"fmt"
	"testing"
)

func TestGradebook(t *testing.T) {
	var tests = []struct {
		a, b, c int
		want    rune
	}{
		{95, 90, 93, 'A'},
		{100, 85, 96, 'A'},
		{92, 93, 94, 'A'},
		{70, 70, 100, 'B'},
		{82, 85, 87, 'B'},
		{84, 79, 85, 'B'},
		{89, 89, 90, 'B'},
		{70, 70, 70, 'C'},
		{75, 70, 79, 'C'},
		{60, 82, 76, 'C'},
		{65, 70, 59, 'D'},
		{66, 62, 68, 'D'},
		{58, 62, 70, 'D'},
		{44, 55, 52, 'F'},
		{48, 55, 52, 'F'},
		{58, 59, 60, 'F'},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d, %d", test.a, test.b, test.c)
		t.Run(name, func(t *testing.T) {
			grade := GetGrade(test.a, test.b, test.c)

			if grade != test.want {
				t.Errorf("GetGrade(%d, %d, %d) = %c; want %c", test.a, test.b, test.c, grade, test.want)
			}
		})
	}
}
