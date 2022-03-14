package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := os.Stdin

	fmt.Print("What is your current age? ")
	age := inputToInt(input)

	fmt.Print("At what age would you like to retire? ")
	retirementAge := inputToInt(input)

	yearsLeft := retirementAge - age
	fmt.Printf("You have %d years until you can retire.\n", yearsLeft)

	currentTime := time.Now()

	retirementYear := currentTime.Year() + yearsLeft

	fmt.Printf("It's %d, so you can retire in %d\n", currentTime.Year(), retirementYear)
}

func inputToInt(input io.Reader) int {
	reader := bufio.NewReader(input)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	value, _ := strconv.ParseInt(userInput, 10, 0)

	return int(value)
}
