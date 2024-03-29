package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	support "support/go/support.go/support/go"
)

func main() {
	input := os.Stdin
	reader := bufio.NewReader(input)

	length := floatInputValidation(reader, "What is the length of the room in feet? ")

	width := floatInputValidation(reader, "What is the width of the room in feet? ")

	fmt.Printf("You entered dimensions of %f feet by %f feet.\n", length, width)

	imperial_area, metric_area := imperialToMetricArea(length, width)

	fmt.Println("The area is:")
	fmt.Printf("%.2f square feet\n", imperial_area)
	fmt.Printf("%.2f square meters\n", metric_area)
}

func imperialToMetricArea(x float64, y float64) (float64, float64) {
	const FEET_TO_METERS = 0.09290304
	imperial_area := x * y
	return imperial_area, imperial_area * FEET_TO_METERS
}

func floatInputValidation(reader *bufio.Reader, question string) float64 {
	var valid_input float64

	for {
		fmt.Print(question)

		entered := support.ReadAndCleanString(reader)
		converted, err := strconv.ParseFloat(entered, 64)

		if err == nil && converted >= 0 {
			valid_input = converted
			break
		}

		fmt.Print("Not a valid number! ")
	}

	return valid_input
}
