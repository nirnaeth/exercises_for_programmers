package main

import "testing"

func TestSum(t *testing.T) {
	first := int64(10)
	second := int64(5)
	expected := int64(15)

	result := Sum(first, second)

	if result != expected {
		t.Errorf("Basic math is not an opinion. Expected %d, got %d\n", expected, result)
	}
}

func TestSub(t *testing.T) {
	first := int64(10)
	second := int64(5)
	expected := int64(5)

	result := Sub(first, second)

	if result != expected {
		t.Errorf("Basic math is not an opinion. Expected %d, got %d\n", expected, result)
	}
}

func TestProduct(t *testing.T) {
	first := int64(10)
	second := int64(5)
	expected := int64(50)

	result := Product(first, second)

	if result != expected {
		t.Errorf("Basic math is not an opinion. Expected %d, got %d\n", expected, result)
	}
}

func TestDivision(t *testing.T) {
	first := int64(10)
	second := int64(5)
	expected := int64(2)

	result := Division(first, second)

	if result != expected {
		t.Errorf("Basic math is not an opinion. Expected %d, got %d\n", expected, result)
	}
}
