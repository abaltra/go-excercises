package main

import "testing"

func TestValidParenthesisInvalid(t *testing.T) {
	input := "(]"
	result := isValid(input)

	if result == true {
		t.Errorf("Expected '%s' to be false, got true", input)
	}
}

func TestValidParenthesisValid(t *testing.T) {
	input := "()"
	result := isValid(input)

	if result == false {
		t.Errorf("Expected '%s' to be true, got false", input)
	}
}

func TestValidParenthesisEmpty(t *testing.T) {
	input := ""
	result := isValid(input)

	if result == false {
		t.Errorf("Expected '%s' to be true, got false", input)
	}
}

func TestValidParenthesisNonBalanced(t *testing.T) {
	input := "()["
	result := isValid(input)

	if result == true {
		t.Errorf("Expected '%s' to be false, got true", input)
	}
}
