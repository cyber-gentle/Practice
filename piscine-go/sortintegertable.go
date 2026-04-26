package main

import "fmt"

func SortIntegerTable(table []int) {
	var result []int

	for i := len(table)-1; i >= 0; i-- {
		result = append(result, table[i])
	}
	fmt.Println(result)
}