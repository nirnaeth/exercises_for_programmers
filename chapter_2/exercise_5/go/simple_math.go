package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Stdin
	reader := bufio.NewReader(input)

	fmt.Print("Enter the first number: ")
	input_1, _ := reader.ReadString('\n')
	input_1 = strings.TrimSpace(input_1)
	number_1, _ := strconv.ParseInt(input_1, 10, 64)

	fmt.Print("Enter the first number: ")
	input_2, _ := reader.ReadString('\n')
	input_2 = strings.TrimSpace(input_2)
	number_2, _ := strconv.ParseInt(input_2, 10, 64)

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
