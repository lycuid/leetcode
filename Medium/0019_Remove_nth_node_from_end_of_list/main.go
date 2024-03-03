// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(head *ListNode, depth int, n int) (*ListNode, int) {
	if head == nil {
		return head, depth - 1
	}
	new_head, size := Aux(head.Next, depth+1, n)
	if size-depth == n {
		return new_head, size
	}
	head.Next = new_head
	return head, size
}

func removeNthFromEnd(head *ListNode, n int) (new_head *ListNode) {
	new_head, _ = Aux(head, 1, n-1)
	return new_head
}

func main() {}
