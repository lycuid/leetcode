// https://leetcode.com/problems/delete-node-in-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}

func main() {}
