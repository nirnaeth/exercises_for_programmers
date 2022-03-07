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

	fmt.Print("What is the quote? ")
	input_quote, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
	input_quote = strings.TrimSpace(input_quote)

	fmt.Print("Who said it? ")
	input_author, _ := reader.ReadString('\n') // ReadString will block until the delimiter is entered
	input_author = strings.TrimSpace(input_author)

	quote := input_author + " says, \"" + input_quote + "\"\n"

	fmt.Print(quote)
}
