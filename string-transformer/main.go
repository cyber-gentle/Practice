package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"transform/gen2"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	command string
)

func main() {
	fmt.Println(" SENTINEL STRING TRANSFORMER — ONLINE \n ──────────────────────────────────────")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input_slice := strings.Fields(strings.TrimSpace(input))

		command = input_slice[0]

		switch gen2.LowerCase(command) {
		case "upper":
			transformed_word := gen2.UpperCase(strings.Join(input_slice[1:], " "))
			fmt.Printf("→ %v\n\n", transformed_word)

		case "lower":
			transformed_word := gen2.LowerCase(strings.Join(input_slice[1:], " "))
			fmt.Printf("→ %v\n\n", transformed_word)

		case "cap":
			transformed_word := gen2.Capitalize(strings.Join(input_slice[1:], " "))
			fmt.Printf("→ %v\n\n", transformed_word)

		case "title":

		case "snake":

		case "reverse":

		case "exit":
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}
	}

}
