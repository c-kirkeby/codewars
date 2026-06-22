package main

func GetGrade(a, b, c int) rune {
	score := (a + b + c) / 3
	switch {
	case score < 60:
		return 'F'
	case score < 70:
		return 'D'
	case score < 80:
		return 'C'
	case score < 90:
		return 'B'
	default:
		return 'A'
	}
}
