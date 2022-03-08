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

	fmt.Print("Enter a noun: ")
	noun, _ := reader.ReadString('\n')
	noun = strings.TrimSpace(noun)

	fmt.Print("Enter a verb: ")
	verb, _ := reader.ReadString('\n')
	verb = strings.TrimSpace(verb)

	fmt.Print("Enter an adjective: ")
	adjective, _ := reader.ReadString('\n')
	adjective = strings.TrimSpace(adjective)

	fmt.Print("Enter an adverb: ")
	adverb, _ := reader.ReadString('\n')
	adverb = strings.TrimSpace(adverb)

	output := fmt.Sprintf("Do you %s your %s %s %s? That's hilarious!", verb, adjective, noun, adverb)

	fmt.Println(output)
}
