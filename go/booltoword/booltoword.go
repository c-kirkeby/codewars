package booltoword

func BoolToWord(word bool) (result string) {
	switch word {
	case true:
		result = "Yes"
	case false:
		result = "No"
	}
	return result
}
