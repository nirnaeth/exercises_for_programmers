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

	fmt.Print("What is the length of the room in feet? ")
	// can be substituted with ReadAndCleanString(input)
	reader := bufio.NewReader(input)
	input_length := support.ReadAndCleanString(reader)
	length, _ := strconv.ParseInt(input_length, 10, 64)

	fmt.Print("What is the width of the room in feet? ")
	// can be substituted with ReadAndCleanString(input)
	input_width := support.ReadAndCleanString(reader)
	width, _ := strconv.ParseInt(input_width, 10, 64)

	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", length, width)

	const FEET_TO_METERS = 0.09290304
	imperial_area := float64(length * width)
	metric_area := imperial_area * FEET_TO_METERS
	casted_imperial_area := int(imperial_area)

	fmt.Println("The area is:")
	fmt.Printf("%d square feet\n", casted_imperial_area)
	fmt.Printf("%f square meters\n", metric_area)
}
