package printdigits

import (
	"github.com/01-edu/z01"
)

func PrintDigits() {
	for d := '0'; d <= '9'; d++ {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}
