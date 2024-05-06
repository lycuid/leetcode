// https://leetcode.com/problems/remove-nodes-from-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head != nil {
		if head.Next = removeNodes(head.Next); head.Next != nil && head.Next.Val > head.Val {
			return head.Next
		}
	}
	return head
}

func main() {}
