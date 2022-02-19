package main

import "testing"

func TestCalculateTip(t *testing.T) {
	bill := 100.0
	rate := 20.0

	expected := 20.0

	tip := CalculateTip(bill, rate)

	if tip != expected {
		t.Errorf("got %f, expected %f", tip, expected)
	}
}

func TestNumberToFloat(t *testing.T) {
	input := "100"

	expected := 100.0

	number := InputToFloat(input)

	if number != expected {
		t.Errorf("got %f, expected %f", number, expected)
	}
}
