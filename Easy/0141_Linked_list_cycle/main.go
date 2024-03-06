// https://leetcode.com/problems/linked-list-cycle/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return fast == slow
}

func main() {}
