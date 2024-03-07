// https://leetcode.com/problems/middle-of-the-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	if fast.Next != nil {
		slow = slow.Next
	}
	return slow
}

func main() {}
