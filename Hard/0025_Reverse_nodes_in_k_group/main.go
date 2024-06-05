// https://leetcode.com/problems/reverse-nodes-in-k-group/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetLast(head *ListNode, k, curr int) *ListNode {
	if head == nil || curr == k {
		return head
	}
	return GetLast(head.Next, k, curr+1)
}

func Reverse(head, last *ListNode) {
	if head == nil || head == last {
		return
	}
	next := head.Next
	if Reverse(next, last); next != nil {
		next.Next = head
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if last := GetLast(head, k, 1); last != nil {
		next := reverseKGroup(last.Next, k)
		Reverse(head, last)
		head.Next = next
		return last
	}
	return head
}

func main() {}
