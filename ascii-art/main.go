package main

import (
	"ascii-art/processor"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		input      string
		bannerName string
		validated  bool
		err        error
	)

	switch len(os.Args) {
	case 2:
		input = os.Args[1]
		bannerName = "standard.txt"

	case 3:
		input = os.Args[1]
		bannerName = strings.ToLower(os.Args[2])

	default:
		fmt.Println("Usage go run . <text> <bannertype>.")
		return
	}

	bannerMap, err := processor.LoadBanner(bannerName)
	if err != nil {
		fmt.Println(err)
		return
	}

	words := processor.SplitInput(input)

	for _, word := range words {
		validated, err = processor.ValidateInput(word)
		if err != nil || !validated {
			fmt.Println(err)
			return
		} else {
			processor.GenerateArt(word, bannerMap)
		}
	}
}
