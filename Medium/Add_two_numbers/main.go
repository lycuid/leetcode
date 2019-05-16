// https://leetcode.com/problems/add-two-numbers/

package main

// ListNode generic
type ListNode struct {
	Val  int
	Next *ListNode
}

func addListNode(l1, l2 *ListNode, carry int) *ListNode {
	var (
		l      ListNode
		l1Next *ListNode
		l2Next *ListNode
	)

	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	if l1 != nil {
		carry += l1.Val
		if l1.Next != nil {
			l1Next = l1.Next
		}
	}
	if l2 != nil {
		carry += l2.Val
		if l2.Next != nil {
			l2Next = l2.Next
		}
	}

	l = ListNode{
		Val:  carry % 10,
		Next: addListNode(l1Next, l2Next, carry/10),
	}

	return &l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addListNode(l1, l2, 0)
}

func main() {}
