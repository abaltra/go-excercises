package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	result := productExceptSelf(input)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v to equal %v", expected, result)
	}

}

func TestProductExceptSelfWithZeroes(t *testing.T) {
	input := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	result := productExceptSelf(input)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v to equal %v", expected, result)
	}

}
