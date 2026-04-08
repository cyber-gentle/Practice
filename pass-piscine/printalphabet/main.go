package pass_piscine

import (
	"github.com/01-edu/z01"
)

func Printalphabet() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}