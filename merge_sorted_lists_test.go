package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}
	l1.Next.Next = &ListNode{Val: 4}

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 5}

	result := mergeTwoLists(&l1, &l2)
	expected := []int{1, 1, 2, 3, 4, 5}

	if len(expected) != len(result.ToArray()) {
		t.Errorf("Expected result to have length %d, had %d", len(expected), len(result.ToArray()))
	}

	for idx, val := range result.ToArray() {
		if expected[idx] != val {
			t.Errorf("Expected %d to be in index %d, found %d", expected[idx], idx, val)
		}
	}
}

func TestMergeTwoListsEmptyL1(t *testing.T) {
	var l1 *ListNode
	l1 = nil
	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 5}

	result := mergeTwoLists(l1, &l2)
	expected := []int{1, 2, 5}

	if len(expected) != len(result.ToArray()) {
		t.Errorf("Expected result to have length %d, had %d", len(expected), len(result.ToArray()))
	}

	for idx, val := range result.ToArray() {
		if expected[idx] != val {
			t.Errorf("Expected %d to be in index %d, found %d", expected[idx], idx, val)
		}
	}
}

func TestMergeTwoListsEmptyL2(t *testing.T) {
	var l1 *ListNode
	l1 = nil
	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 5}

	result := mergeTwoLists(&l2, l1)
	expected := []int{1, 2, 5}

	if len(expected) != len(result.ToArray()) {
		t.Errorf("Expected result to have length %d, had %d", len(expected), len(result.ToArray()))
	}

	for idx, val := range result.ToArray() {
		if expected[idx] != val {
			t.Errorf("Expected %d to be in index %d, found %d", expected[idx], idx, val)
		}
	}
}

func TestMergeTwoListsEmptyBoth(t *testing.T) {
	var l1 *ListNode
	var l2 *ListNode
	l1 = nil
	l2 = nil

	result := mergeTwoLists(l1, l2)
	expected := []int{}

	if len(expected) != len(result.ToArray()) {
		t.Errorf("Expected result to have length %d, had %d", len(expected), len(result.ToArray()))
	}

	for idx, val := range result.ToArray() {
		if expected[idx] != val {
			t.Errorf("Expected %d to be in index %d, found %d", expected[idx], idx, val)
		}
	}
}
