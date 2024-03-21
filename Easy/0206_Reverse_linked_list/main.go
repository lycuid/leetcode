// https://leetcode.com/problems/reverse-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) (new_head *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}
	new_head = reverseList(head.Next)
	head.Next.Next, head.Next = head, nil
	return new_head
}

func main() {}
