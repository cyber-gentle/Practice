package pass_piscine

import (
	"github.com/01-edu/z01"
)

func Printdigits() {
	for d := '0'; d <= '9'; d++ {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}