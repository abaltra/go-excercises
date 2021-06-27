package main

/*
https://leetcode.com/problems/merge-two-sorted-lists/

Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.

Input: l1 = [1,2,4], l2 = [1,3,4]
Output: [1,1,2,3,4,4]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) ToArray() []int {
	res := []int{}

	runner := n

	for runner != nil {
		res = append(res, runner.Val)
		runner = runner.Next
	}

	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	adder := &dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			adder.Next = l1
			adder = adder.Next
			l1 = l1.Next
		} else {
			adder.Next = l2
			adder = adder.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		adder.Next = l1
	}

	if l2 != nil {
		adder.Next = l2
	}

	return dummy.Next
}
