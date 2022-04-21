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

	var length float64
	var width float64

	for {
		fmt.Print("What is the length of the room in feet? ")

		entered := support.ReadAndCleanString(reader)
		converted, err := strconv.ParseFloat(entered, 64)

		if err == nil && converted >= 0 {
			length = converted
			break
		}

		fmt.Print("Not a valid number! ")
	}

	for {
		fmt.Print("What is the width of the room in feet? ")

		entered := support.ReadAndCleanString(reader)
		converted, err := strconv.ParseFloat(entered, 64)

		if err == nil && converted >= 0 {
			width = converted
			break
		}

		fmt.Print("Not a valid number! ")
	}

	fmt.Printf("You entered dimensions of %f feet by %f feet.\n", length, width)

	const FEET_TO_METERS = 0.09290304
	imperial_area := float64(length * width)
	metric_area := imperial_area * FEET_TO_METERS
	casted_imperial_area := int(imperial_area)

	fmt.Println("The area is:")
	fmt.Printf("%d square feet\n", casted_imperial_area)
	fmt.Printf("%f square meters\n", metric_area)
}
