package main

import (
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

	if os.Getenv("CRASH") == "1" {
		InputToFloat(invalid_bill)
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
