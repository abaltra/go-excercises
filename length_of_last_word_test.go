package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	lastWord := "world"
	input := fmt.Sprintf("Hello %s", lastWord)
	output := lengthOfLastWord(input)

	if output != len(lastWord) {
		t.Errorf("Expected %d to be %d", output, len(lastWord))
	}
}

func TestLengthOfLastWordEmptyLast(t *testing.T) {
	lastWord := " "
	input := fmt.Sprintf("Hello %s", lastWord)
	output := lengthOfLastWord(input)

	if output != 5 {
		t.Errorf("Expected %d to be %d", output, 5)
	}
}

func TestLengthOfLastWordEmtpyString(t *testing.T) {
	input := fmt.Sprintf("")
	output := lengthOfLastWord(input)

	if output != 0 {
		t.Errorf("Expected %d to be %d", output, 0)
	}
}

func TestLengthOfLastWordWordWithSpaces(t *testing.T) {
	lastWord := "world "
	input := lastWord
	output := lengthOfLastWord(input)

	if output != len(strings.Trim(lastWord, " ")) {
		t.Errorf("Expected %d to be %d", output, len(strings.Trim(input, " ")))
	}
}
