// https://leetcode.com/problems/swap-nodes-in-pairs/

package main

// ListNode the definition for listnode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(head.Next.Next)
	next.Next = head
	return next
}

func main() { }


