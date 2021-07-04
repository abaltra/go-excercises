package main

import "testing"

func TestZigZagConversion(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 3
	expected := "PAHNAPLSIIGYIR"

	result := convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}

func TestZigZagConversion2(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 4
	expected := "PINALSIGYAHRPI"

	result := convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}

func TestZigZagConversionOneRow(t *testing.T) {
	input := "PAYPALISHIRING"
	numRows := 1
	expected := "PAYPALISHIRING"

	result := convert(input, numRows)

	if result != expected {
		t.Errorf("Expected %s to equal %s", result, expected)
	}
}
