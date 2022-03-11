package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var number_1 int64
	var number_2 int64

	input := os.Stdin
	reader := bufio.NewReader(input)

	for {
		fmt.Print("Enter the first number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseInt(input, 10, 64)

		if err == nil && value >= 0 {
			number_1 = value
			break
		}

		fmt.Print("Not a valid number! ")
	}

	for {
		fmt.Print("Enter the second number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseInt(input, 10, 64)

		if err == nil && value >= 0 {
			number_2 = value
			break
		}

		fmt.Print("Not a valid number! ")
	}

	fmt.Printf("%d + %d = %d\n", number_1, number_2, Sum(number_1, number_2))
	fmt.Printf("%d - %d = %d\n", number_1, number_2, Sub(number_1, number_2))
	fmt.Printf("%d * %d = %d\n", number_1, number_2, Product(number_1, number_2))
	fmt.Printf("%d / %d = %d\n", number_1, number_2, Division(number_1, number_2))
}

func Sum(first int64, second int64) int64 {
	return first + second
}

func Sub(first int64, second int64) int64 {
	return first - second
}

func Product(first int64, second int64) int64 {
	return first * second
}

func Division(first int64, second int64) int64 {
	return first / second
}
