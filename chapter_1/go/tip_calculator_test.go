package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestCalculateTip(t *testing.T) {
	bill := 100.0
	rate := 20.0

	expected := 20.0

	tip := CalculateTip(bill, rate)

	if tip != expected {
		t.Errorf("got %f, expected %f", tip, expected)
	}
}

// Code from https://stackoverflow.com/questions/26225513/how-to-test-os-exit-scenarios-in-go
func TestExitWithErrorWhenInputIsNotANumber(t *testing.T) {
	invalid_bill := "not a number"

	if os.Getenv("CRASH") == "1" { // setting env to test the wrongful exit status
		output := captureOutput(func() { // capturing output to verify message, see support function at the bottom
			InputToFloat(invalid_bill)
		})

		expected_message := "Amount entered 'not a number' is not a number\n"

		if output != expected_message { // this is actually not working
			t.Errorf("got %s, expected %s", output, expected_message)
		}
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExitWithErrorWhenInputIsNotANumber")
	cmd.Env = append(os.Environ(), "CRASH=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestNumberToFloat(t *testing.T) {
	input := "100"

	expected := 100.0

	number := InputToFloat(input)

	if number != expected {
		t.Errorf("got %f, expected %f", number, expected)
	}
}

// Support
// captures the output to stdout to test the messages
func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}
