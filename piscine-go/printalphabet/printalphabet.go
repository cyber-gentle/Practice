package printalphabet

import (
	"github.com/01-edu/z01"
)

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
