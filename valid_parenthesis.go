package main

/*
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

type stack []rune

func (s *stack) peek() (rune, bool) {
	idx := len(*s) - 1
	if idx < 0 {
		return '_', false
	}

	return (*s)[idx], true
}

func (s *stack) push(n rune) {
	*s = append(*s, n)
}

func (s *stack) pop() (rune, bool) {
	l := len(*s) - 1
	if l < 0 {
		return -1, true
	}

	elem := (*s)[l]
	*s = (*s)[:l]

	return elem, false
}

func isValidPair(opener rune, closer rune) bool {
	if opener == '{' && closer != '}' {
		return false
	}

	if opener == '[' && closer != ']' {
		return false
	}

	if opener == '(' && closer != ')' {
		return false
	}

	return true
}

func isValid(s string) bool {
	var _s stack
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			_s.push(c)
		} else if c == '}' || c == ']' || c == ')' {
			opener, empty := _s.pop()
			if empty {
				return false
			}

			if !isValidPair(opener, c) {
				return false
			}
		}
	}

	if len(_s) > 0 {
		return false
	}

	return true
}
