package processor

import "fmt"

func ValidateInput(word string) (bool, error) {
	var (
		validated bool = true
	)

	for i, char := range word {
		if char < 32 || char > 126 {
			return !validated, fmt.Errorf(
				"char at index %d is not a valid character", i)
		}
	}
	return validated, nil
}
