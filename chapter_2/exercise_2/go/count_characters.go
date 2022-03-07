// Create a program that prompts for an input string and displays the input string
// and the number of characters the string contains.
// The program has been extended with the challenges.

package main

import (
	"fmt"
	"os"
	support "support/go/support.go/support/go"
)

func main() {
	input := os.Stdin

	fmt.Println("What is the input string?")

	// reader := bufio.NewReader(input) // Read from stdin
	// string, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
	// string = strings.TrimSpace(string)   // remove the delimeter from the string

	word := support.ReadAndCleanString(input)
	length := len(word)

	fmt.Printf("%s has %d characters\n", word, length)
}
