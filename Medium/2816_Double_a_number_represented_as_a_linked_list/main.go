// https://leetcode.com/problems/double-a-number-represented-as-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(node *ListNode) {
	if node != nil {
		Aux(node.Next)
		if node.Val *= 2; node.Next != nil && node.Next.Val >= 10 {
			node.Val += node.Next.Val / 10
			node.Next.Val %= 10
		}
	}
}

func doubleIt(head *ListNode) *ListNode {
	for Aux(head); head != nil && head.Val >= 10; {
		head = &ListNode{
			Next: head,
			Val:  head.Val / 10,
		}
		head.Next.Val %= 10
	}
	return head
}

func main() {}
