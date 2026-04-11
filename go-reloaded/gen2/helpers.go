package gen2

import (
	"strings"
)

func UpperCaseLastWord(words []string) string {
	lastWord := strings.ToUpper(words[0])
	return lastWord
}

func LowerCaseLastWord(words []string) string {
	lastWord := strings.ToLower(words[0])
	return lastWord
}

func CapitaliseLastWord(words []string) string {
	lastWord := strings.ToUpper(words[0][:1]) + strings.ToLower(words[0][1:])
	return lastWord
}

func UpperCaseLastWordN(words []string) string {
	
	lastWordN := strings.ToUpper(strings.Join(words, " "))
	return lastWordN
}

func LowerCaseLastWordN(words []string) string {
	
	lastWordN := strings.ToLower(strings.Join(words, " "))
	return lastWordN
}

func CapitaliseLastWordN(words []string) string {

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToUpper(string(words[i][0])) + strings.ToLower(string(words[i][1:]))
	}
	
	lastWordN := strings.Join(words, " ")
	return lastWordN
}