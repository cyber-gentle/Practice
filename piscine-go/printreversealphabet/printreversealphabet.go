package printreversealphabet

import (
	"github.com/01-edu/z01"
)

func PrintReverseAlphabet() {
	for c := 'z'; c >= 'a'; c-- {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
