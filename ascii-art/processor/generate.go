package processor

import (
	"fmt"
)

func GenerateArt(word string, bannerMap map[rune][]string) {
	if word == "" || len(word) == 0 {
		fmt.Println()
	} else {
		Render(word, bannerMap)
	}
}
