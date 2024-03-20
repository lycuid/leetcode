// https://leetcode.com/problems/merge-in-between-linked-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := list1
	for ; a > 1; a, b = a-1, b-1 {
		head = head.Next
	}
	head.Next, head = list2, head.Next
	for list2.Next != nil {
		list2 = list2.Next
	}
	for ; b > 0; b-- {
		head = head.Next
	}
	list2.Next = head
	return list1
}

func main() {}
