// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func At(l *ListNode, i uint) *ListNode {
	if i == 0 || l == nil {
		return l
	}
	return At(l.Next, i-1)
}

func swap(head *ListNode, x, y uint) *ListNode {
	lx, ly := At(head, x), At(head, y)
	lxp, lyp := head, head
	for ; lxp != nil && lxp.Next != lx; lxp = lxp.Next {
	}
	for ; lyp != nil && lyp.Next != ly; lyp = lyp.Next {
	}
	if lxp != nil {
		lxp.Next = ly
	} else {
		head = ly
	}
	if lyp != nil {
		lyp.Next = lx
	} else {
		head = lx
	}
	lx.Next, ly.Next = ly.Next, lx.Next
	return head
}

func swapNodes(head *ListNode, k int) *ListNode {
	count := 0
	for l := head; l != nil; l = l.Next {
		count++
	}
	return swap(head, uint(k-1), uint(count-k))
}

func main() {}
