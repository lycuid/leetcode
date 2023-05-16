// https://leetcode.com/problems/swap-nodes-in-pairs/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) (next *ListNode) {
	if head != nil {
		if next = head.Next; next != nil {
			head.Next, next.Next = swapPairs(next.Next), head
		} else {
			next = head
		}
	}
	return next
}

func main() {}
