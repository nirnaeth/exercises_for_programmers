/* Create a simple tip calculator.
The program should prompt for a bill amount and a tip rate.
The program must compute the tip and display both the tip and the total
amount of the bill.

The program has been extended with the challenges.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Read from stdin

	fmt.Print("Enter the bill amount: ")
	input, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
	bill := InputToFloat(input)

	fmt.Print("Enter the rate: ")
	input, _ = reader.ReadString('\n') // ReadString will block until the delimiter is entered
	rate := InputToFloat(input)

	tip := CalculateTip(bill, rate)
	total := bill + tip

	fmt.Println("The tip is ", tip)
	fmt.Println("The total is ", total)
}

func CalculateTip(bill float64, rate float64) float64 {
	return math.Round(bill / 100 * rate)
}

func InputToFloat(input string) float64 {
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string

	converted, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Printf("Amount entered '%s' is not a number\n", input)
		os.Exit(1)
	}

	return converted
}
