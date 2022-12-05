// https://leetcode.com/problems/middle-of-the-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mid, node, step := head, head, true
	for node != nil {
		if node, step = node.Next, !step; step {
			mid = mid.Next
		}
	}
	return mid
}

func main() {}
