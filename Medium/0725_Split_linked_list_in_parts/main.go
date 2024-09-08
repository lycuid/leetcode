// https://leetcode.com/problems/split-linked-list-in-parts/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	nodes, n := make([]*ListNode, k), 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	for i := range nodes {
		nodes[i] = &ListNode{Val: n / k}
	}
	for i := 0; i < n%k; i++ {
		nodes[i].Val++
	}
	for i, tmp := range nodes {
		nodes[i] = head
		for j := 1; j < tmp.Val; j++ {
			head = head.Next
		}
		if head != nil {
			head, head.Next = head.Next, nil
		}
	}
	return nodes
}

func main() {}
