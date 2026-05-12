package processor

import (
	"fmt"
)

func Render(word string, bannerMap map[rune][]string) {
	for row := range 8 {
		for _, char := range word {
			fmt.Print(bannerMap[char][row])
		}
		fmt.Println()
	}



}
