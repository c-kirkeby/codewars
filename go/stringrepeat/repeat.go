package main

import "strings"

func RepeatStr(repetitions int, value string) string {
	var out strings.Builder
	for i := 0; i <= repetitions; i++ {
		out.WriteString(value)
	}
	return out.String()
}
