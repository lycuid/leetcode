// https://leetcode.com/problems/intersection-of-two-linked-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(a, b *ListNode) *ListNode {
	var size int
	for head := a; head != nil; head = head.Next {
		size++
	}
	for head := b; head != nil; head = head.Next {
		size--
	}
	for ; size > 0; size-- {
		a = a.Next
	}
	for ; size < 0; size++ {
		b = b.Next
	}
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a, b = a.Next, b.Next
	}
	return nil
}

func main() {}
