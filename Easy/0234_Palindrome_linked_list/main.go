// https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(list *ListNode, rev **ListNode) bool {
	if list == nil {
		return true
	}
	if Aux(list.Next, rev) && list.Val == (*rev).Val {
		*rev = (*rev).Next
		return true
	}
	return false
}

func isPalindrome(head *ListNode) bool {
	return Aux(head, &head)
}

func main() {}
