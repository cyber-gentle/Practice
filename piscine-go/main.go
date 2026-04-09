package main

import (
	"fmt"
	//"github.com/01-edu/z01"
)

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}
