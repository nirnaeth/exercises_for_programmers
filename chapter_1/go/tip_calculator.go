/* Create a simple tip calculator.
The program should prompt for a bill amount and a tip rate.
The program must compute the tip and display both the tip and the total
amount of the bill.
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
	bill := input_to_float(reader)

	fmt.Print("Enter the rate: ")
	rate := input_to_float(reader)

	tip := calculateTip(bill, rate)
	total := bill + tip

	fmt.Println("The tip is ", tip)
	fmt.Println("The total is ", total)
}

func calculateTip(bill float64, rate float64) float64 {
	return math.Round(bill / 100 * rate)
}

func input_to_float(reader *bufio.Reader) float64 {
	input, _ := reader.ReadString('\n')     // ReadString will block until the delimiter is entered
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string
	input_in_float, _ := strconv.ParseFloat(input, 64)

	return input_in_float
}
