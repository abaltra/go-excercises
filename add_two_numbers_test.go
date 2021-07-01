package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	a := ListNode{Val: 2}
	a.Next = &ListNode{Val: 4}
	a.Next.Next = &ListNode{Val: 3}

	b := ListNode{Val: 5}
	b.Next = &ListNode{Val: 6}
	b.Next.Next = &ListNode{Val: 4}

	expected := &ListNode{Val: 7}
	expected.Next = &ListNode{Val: 0}
	expected.Next.Next = &ListNode{Val: 8}

	result := addTwoNumbers(&a, &b)

	for result != nil && expected != nil {
		if result.Val != expected.Val {
			t.Errorf("Expected %d to equal %d", result.Val, expected.Val)
		}
		result = result.Next
		expected = expected.Next
	}

	if result != nil || expected != nil {
		t.Error("Expected result to have length of 3")
	}
}

func TestAddTwoNumbersUnequalNumberLengths(t *testing.T) {
	a := ListNode{Val: 2}
	a.Next = &ListNode{Val: 4}
	a.Next.Next = &ListNode{Val: 3}

	b := ListNode{Val: 5}
	b.Next = &ListNode{Val: 6}

	expected := &ListNode{Val: 7}
	expected.Next = &ListNode{Val: 0}
	expected.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(&a, &b)

	for result != nil && expected != nil {
		if result.Val != expected.Val {
			t.Errorf("Expected %d to equal %d", result.Val, expected.Val)
		}
		result = result.Next
		expected = expected.Next
	}

	if result != nil || expected != nil {
		t.Error("Expected result to have length of 3")
	}
}
