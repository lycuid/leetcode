// https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Tuple struct{ parent, child *ListNode }

func Get(head *ListNode, k int) (ret Tuple) {
	for ret.parent, ret.child = nil, head; k > 1; k-- {
		ret.parent, ret.child = ret.child, ret.child.Next
	}
	return ret
}

func Swap(head *ListNode, p1, p2 Tuple) *ListNode {
	if p1.parent != nil {
		p1.parent.Next = p2.child
	} else {
		head = p2.child
	}
	if p2.parent != nil {
		p2.parent.Next = p1.child
	} else {
		head = p2.child
	}
	p1.child.Next, p2.child.Next = p2.child.Next, p1.child.Next
	return head
}

func swapNodes(head *ListNode, k int) *ListNode {
	count := 1
	for root := head; root != nil; root = root.Next {
		count++
	}
	if k > count/2 {
		return Swap(head, Get(head, count-k), Get(head, k))
	}
	return Swap(head, Get(head, k), Get(head, count-k))
}

func main() {}
