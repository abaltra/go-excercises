package main

import "testing"

func TestNumeralToRoman(t *testing.T) {
	input := 123
	expected := "CXXIII"

	result := intToRoman(input)

	if expected != result {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}

func TestNumeralToRomanBase(t *testing.T) {
	input := 1
	expected := "I"

	result := intToRoman(input)

	if expected != result {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}

func TestNumeralToRomanFours(t *testing.T) {
	input := 4
	expected := "IV"

	result := intToRoman(input)

	if expected != result {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}
