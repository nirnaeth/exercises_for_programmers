package support

import (
	"bufio"
	"io"
	"strings"
)

func ReadAndCleanString(input io.Reader) string {
	reader := bufio.NewReader(input)                     // Read from stdin
	original_string, _ := reader.ReadString('\n')        // ReadString will block until the delimiter is entered
	cleaned_string := strings.TrimSpace(original_string) // remove the delimeter from the string

	return cleaned_string
}
