// https://leetcode.com/problems/reorder-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(right *ListNode, left **ListNode) {
	if right != nil {
		Aux(right.Next, left)
		next := (*left).Next
		if (*left).Next = right; *left == right || next == right {
			right.Next = nil
		} else {
			right.Next = next
		}
		*left = next
	}
}

func reorderList(head *ListNode) {
	if head != nil {
		left, right := head, head
		for node, step := head, false; node.Next != nil; node = node.Next {
			if step = !step; step {
				right = right.Next
			}
		}
		Aux(right, &left)
	}
}

func main() {}
