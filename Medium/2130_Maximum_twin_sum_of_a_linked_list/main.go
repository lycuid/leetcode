// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func solve(head, rhead *ListNode) (*ListNode, int) {
	if rhead == nil {
		return head, 0
	}
	node, val := solve(head, rhead.Next)
	return node.Next, max(val, node.Val+rhead.Val)
}

func pairSum(head *ListNode) int {
	_, val := solve(head, head)
	return val
}

func main() {}
