package main

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
func addNums(a int, b int, c int) (int, int) {
	r := a + b + c
	carry := 0
	if r > 9 {
		carry = 1
		r = r % 10
	}

	return r, carry
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	result := dummy
	carry := 0
	sum := 0

	for l1 != nil || l2 != nil {
		a := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		b := 0
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum, carry = addNums(a, b, carry)

		result.Next = &ListNode{Val: sum}
		result = result.Next
	}

	if carry == 1 {
		result.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}
