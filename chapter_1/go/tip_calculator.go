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
	fmt.Print("Enter the bill amount: ")
	reader := bufio.NewReader(os.Stdin)     // Read from stdin
	input, _ := reader.ReadString('\n')     // ReadString will block until the delimiter is entered
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string
	bill, _ := strconv.ParseFloat(input, 64)

	fmt.Print("Enter the rate: ")
	reader = bufio.NewReader(os.Stdin)      // Read from stdin
	input, _ = reader.ReadString('\n')      // ReadString will block until the delimiter is entered
	input = strings.TrimSuffix(input, "\n") // remove the delimeter from the string
	rate, _ := strconv.ParseFloat(input, 64)

	tip := calculateTip(bill, rate)
	total := bill + tip

	fmt.Println("The tip is ", tip)
	fmt.Println("The total is ", total)
}

func calculateTip(bill float64, rate float64) float64 {
	return math.Round(bill / 100 * rate)
}
