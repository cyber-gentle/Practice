package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(`Enter a word: `)
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	if word == "" {
		panic("Empty string")
	}

	fmt.Print("Choose an operation to execute.\n 1. Last Character \n 2. Capitalize\n 3. Delete an index\n 4. Exit\n ")
	fmt.Print("Select operation to continue: ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)
	c_op, _ := strconv.Atoi(op)

	switch c_op {
	case 1:
		result := string(word[len(word)-1])
		fmt.Printf("The last letter of the word '%v' is '%v'\n", word, result)

	case 2:
		result := strings.ToUpper(word)
		fmt.Printf("The uppercase form of '%v' is '%v'.\n", word, result)

	case 3:
		fmt.Print("Select an index to be deleted: ")
		index, _ := reader.ReadString('\n')
		index = strings.TrimSpace(index)
		c_index, _ := strconv.Atoi(index)

		if c_index >= len(word) || c_index < 0 {
			fmt.Println("Index is out of range")
		}
		result := word[:c_index] + word[c_index+1:]
		fmt.Printf("The word %q without index [%d] is %v\n", word, c_index, result)

	case 4:
		return
	}

}
