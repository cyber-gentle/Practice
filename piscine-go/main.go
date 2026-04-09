package main

import (
	"fmt"
	//"github.com/01-edu/z01"
)

func main() {
a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
