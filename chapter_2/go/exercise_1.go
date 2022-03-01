// Create a program that prompts for your name and prints a greeting using your name
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

	fmt.Println("What is your name? ")
	name, _ := reader.ReadString('\n')    // ReadString will block until the delimiter is entered
	name = strings.TrimSuffix(name, "\n") // remove the delimeter from the string

	message := "Hello, " + name + ", nice to meet you!"
	fmt.Println(message)
}
