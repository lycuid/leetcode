// https://leetcode.com/problems/odd-even-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(head *ListNode) (next *ListNode) {
	if next = head.Next; next != nil {
		head.Next = Aux(next)
	}
	return next
}

func oddEvenList(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		last, snd := head, Aux(head)
		for ; last.Next != nil; last = last.Next {
		}
		last.Next = snd
	}
	return head
}

func main() {}
