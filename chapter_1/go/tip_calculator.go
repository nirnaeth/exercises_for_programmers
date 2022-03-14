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
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	conf := "prod"

	CalculateTotalWithTip(conf)
}

func CalculateTotalWithTip(conf string) {
	var bill, rate float64

	bill = retrieveBill(os.Stdin, bill, conf)
	rate = retrieveRate(os.Stdin, rate, conf)

	tip := calculateTip(bill, rate)
	total := bill + tip

	fmt.Println("The tip is ", tip)
	fmt.Println("The total bill is ", total)
}

func retrieveRate(input io.Reader, rate float64, config string) float64 {
	reader := bufio.NewReader(input) // Read from stdin

	for {
		fmt.Print("Enter the rate: ")
		input, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		value, err := inputToFloat(input)

		if err == nil && value >= 0 {
			rate = value
			break
		}

		fmt.Print("Not a valid number! ")

		if config == "test" {
			return -1
		}
	}

	return rate
}

func retrieveBill(input io.Reader, bill float64, conf string) float64 {
	reader := bufio.NewReader(input) // Read from stdin

	for {
		fmt.Print("Enter the bill amount: ")

		input, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
		value, err := inputToFloat(input)

		if err == nil && value >= 0 {
			bill = value
			break
		}

		fmt.Print("Not a valid number! ")

		if conf == "test" {
			break
		}
	}

	return bill
}

func calculateTip(bill float64, rate float64) float64 {
	return math.Round(bill / 100 * rate)
}

func inputToFloat(input string) (float64, error) {
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string

	converted, err := strconv.ParseFloat(input, 64)

	return converted, err
}
