// https://leetcode.com/problems/linked-list-cycle-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	for seen := make(map[*ListNode]bool); head != nil && !seen[head]; head = head.Next {
		seen[head] = true
	}
	return head
}

func main() {}
