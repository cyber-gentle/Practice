package main

import (

)

func Swap(a *int, b *int) {
	raw_a := *a; raw_b := *b
	*a = raw_b; *b = raw_a
}