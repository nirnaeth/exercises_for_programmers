package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCalculateTip(t *testing.T) {
	bill := 100.0
	rate := 20.0

	expected := 20.0

	tip := calculateTip(bill, rate)

	if tip != expected {
		t.Errorf("got %f, expected %f", tip, expected)
	}
}

// Code from https://stackoverflow.com/questions/26225513/how-to-test-os-exit-scenarios-in-go
// func TestExitWithErrorWhenInputIsNotANumber(t *testing.T) {
// 	invalid_bill := "not a number"

// 	if os.Getenv("CRASH") == "1" { // setting env to test the wrongful exit status
// 		output := captureOutput(func() { // capturing output to verify message, see support function at the bottom
// 			InputToFloat(invalid_bill)
// 		})

// 		expected_message := "Amount entered 'not a number' is not a number\n"

// 		if output != expected_message { // this is actually not working
// 			t.Errorf("got %s, expected %s", output, expected_message)
// 		}
// 		return
// 	}

// 	cmd := exec.Command(os.Args[0], "-test.run=TestExitWithErrorWhenInputIsNotANumber")
// 	cmd.Env = append(os.Environ(), "CRASH=1")
// 	err := cmd.Run()
// 	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
// 		return
// 	}

// 	t.Fatalf("process ran with err %v, want exit status 1", err)
// }

func TestNumberToFloat(t *testing.T) {
	input := "100"

	expected := 100.0

	number, _ := inputToFloat(input)

	if number != expected {
		t.Errorf("got %f, expected %f", number, expected)
	}
}

// func TestRetrieveBillDoesNotAcceptStrings(t *testing.T) {
// 	var bill float64
// 	input_from_keyboard := "test"
// 	config := "test"
// 	reader := strings.NewReader(input_from_keyboard)

// 	output := captureOutput(func() { // capturing output to verify message, see support function at the bottom
// 		retrieveBill(reader, bill, config)
// 	})

// 	expected_message := "Not a valid number!\n"

// 	if output == expected_message {
// 		t.Fatalf("got %s, expected %s", output, expected_message)
// 	}
// }

// func TestRetrieveRateDoesNotAcceptStrings(t *testing.T) {
// 	var rate float64
// 	input_from_keyboard := "test"
// 	config := "test"
// 	reader := strings.NewReader(input_from_keyboard)

// 	output := captureOutput(func() { // capturing output to verify message, see support function at the bottom
// 		retrieveRate(reader, rate, config)
// 	})

// 	expected_message := "Not a valid number!\n"

// 	if output == expected_message {
// 		t.Fatalf("got %s, expected %s", output, expected_message)
// 	}
// }

// Support
// captures the output to stdout to test the messages
func captureOutput(tested_function func()) string {
	var buffer bytes.Buffer
	original := os.Stdout
	reader, writer, _ := os.Pipe()

	tested_function()

	io.Copy(&buffer, reader)

	writer.Close()
	os.Stdout = original

	return buffer.String()
}
