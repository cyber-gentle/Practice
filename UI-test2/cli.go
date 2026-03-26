package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	num1, num2 int
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("CHOOSE AN OPERATION TO EXECUTE\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit\n")
	fmt.Print("Select an operation to continue: ")
	op_c, _ := reader.ReadString('\n')
	op_c = strings.TrimSpace(op_c)
	op, _ := strconv.Atoi(op_c)

	for {
		switch op {
		case 1:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 + num2
			fmt.Printf("%d + %d is = %d\n", num1, num2, result)
		}
		break
	}
}
