// https://leetcode.com/problems/partition-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, val int) *ListNode {
	lhead, rhead := &ListNode{-1, nil}, &ListNode{-1, nil}
	left, right := lhead, rhead
	for head != nil {
		if head.Val < val {
			left.Next, head = head, head.Next
			left = left.Next
		} else {
			right.Next, head = head, head.Next
			right = right.Next
		}
	}
	left.Next, right.Next = rhead.Next, nil
	return lhead.Next
}

func main() {}
