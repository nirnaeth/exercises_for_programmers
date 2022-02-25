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

	var bill, rate float64

	bill = retrieveBill(reader, bill)
	rate = retrieveRate(reader, rate)

	tip := CalculateTip(bill, rate)
	total := bill + tip

	fmt.Println("The tip is ", tip)
	fmt.Println("The total bill is ", total)
}

func retrieveRate(reader *bufio.Reader, rate float64) float64 {
	for {
		fmt.Print("Enter the rate: ")
		input, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		value, err := InputToFloat(input)

		if err == nil {
			rate = value
			break
		}
	}

	return rate
}

func retrieveBill(reader *bufio.Reader, bill float64) float64 {
	for {
		fmt.Print("Enter the bill amount: ")

		input, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		value, err := InputToFloat(input)

		if err == nil {
			bill = value
			break
		}
	}
	return bill
}

func CalculateTip(bill float64, rate float64) float64 {
	return math.Round(bill / 100 * rate)
}

func InputToFloat(input string) (float64, error) {
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string

	converted, err := strconv.ParseFloat(input, 64)

	return converted, err
}
