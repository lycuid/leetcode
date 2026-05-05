// https://leetcode.com/problems/rotate-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head != nil {
		last, l := head, 1
		for ; last.Next != nil; last = last.Next {
			l++
		}
		last.Next = head
		for k = l - k%l; k > 0; k-- {
			last = last.Next
		}
		head, last.Next = last.Next, nil
	}
	return head
}

func main() {}
