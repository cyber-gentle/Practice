package main

import (
	"fmt"
	// "github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))

		//z01.PrintRune(rune(i)) TO BE DONE LATER
	}
	fmt.Println()
}
