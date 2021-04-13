// https://leetcode.com/problems/rotate-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	for h := head; true; h = h.Next {
		length++
		if h.Next == nil {
			h.Next = head
			break
		}
	}
	k %= length

	newhead := head
	for l := 0; l < length-k-1; l++ {
		newhead = newhead.Next
	}
	tmp := newhead
	newhead = newhead.Next
	tmp.Next = nil

	return newhead
}

func main() {}
