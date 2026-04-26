package main

import (
	"strconv"
)

func BasicAtoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}