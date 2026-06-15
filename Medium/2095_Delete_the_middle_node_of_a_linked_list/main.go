// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	var (
		prev *ListNode
		slow = head
	)
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		prev, slow = slow, slow.Next
	}
	if prev == nil {
		return nil
	}
	prev.Next = slow.Next
	return head
}

func main() {}
