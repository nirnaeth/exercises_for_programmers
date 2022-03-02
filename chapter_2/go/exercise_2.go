// Create a program that prompts for an input string and displays the input string
// and the number of characters the string contains.
// The program has been extended with the challenges.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Stdin
	reader := bufio.NewReader(input) // Read from stdin

	fmt.Println("What is the input string?")

	string, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
	string = strings.TrimSpace(string)   // remove the delimeter from the string
	length := len(string)

	fmt.Printf("%s has %d characters\n", string, length)
}

// TODO: Create module for support functions
