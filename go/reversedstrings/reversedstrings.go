package reversedstrings

func Reverse(str string) (result string) {
	for _, value := range str {
		result = string(value) + result
	}
	return result
}
