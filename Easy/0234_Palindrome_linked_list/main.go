// https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(l1 *ListNode, l2 **ListNode) bool {
	if l1 == nil {
		return true
	}
	if Aux(l1.Next, l2) && l1.Val == (*l2).Val {
		(*l2) = (*l2).Next
		return true
	}
	return false
}

func isPalindrome(head *ListNode) bool {
	return Aux(head, &head)
}

func main() {}
