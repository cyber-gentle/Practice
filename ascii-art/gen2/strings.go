package gen2

func UpperCase(s string) string {

	char := []rune(s)
	for i := range char {
		if char[i] >= 90 && char[i] <= 122 {
			char[i] -= 32
		}
	}
	return string(char)
}

func LowerCase(s string) string {

	char := []rune(s)
	for i := range char {
		if char[i] >= 65 && char[i] <= 90 {
			char[i] += 32
		}

	}
	return string(char)
}

func Capitalize(s string) string {
	chars := []rune(s)
	for i := range chars {
		if chars[0] >= 97 && chars[0] <= 122 {
			chars[0] -= 32
		} else if chars[i] == 32 {
			if chars[i+1] >= 90 && chars[i+1] <= 122 {
				chars[i+1] -= 32
			}
		}
	}

	return string(chars)
}
