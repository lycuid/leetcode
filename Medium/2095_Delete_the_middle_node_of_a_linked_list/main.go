// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	prev, slow, fast := (*ListNode)(nil), head, head
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}
	prev.Next = slow.Next
	return head
}

func main() {}
