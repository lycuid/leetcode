// https://leetcode.com/problems/merge-nodes-in-between-zeros/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	for node := head; node != nil; node = node.Next {
		next := node.Next
		for ; next != nil && next.Val != 0; next = next.Next {
			node.Val += next.Val
		}
		node.Next = next.Next
	}
	return head
}

func main() {}
