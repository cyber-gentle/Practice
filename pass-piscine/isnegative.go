package main

import (
	"github.com/01-edu/z01"
)

func isNegative(n int) bool {
	if n < 0 {
		return true
	}
	return false
}

package main

import "piscine"

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}