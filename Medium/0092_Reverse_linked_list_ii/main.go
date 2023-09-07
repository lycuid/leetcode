// https://leetcode.com/problems/reverse-linked-list-ii/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	var pre, p1, p2, post *ListNode
	for p1, right = head, right-left; left > 1; left-- {
		pre, p1 = p1, p1.Next
	}
	for p2, post = p1, p1.Next; right > 0; right-- {
		post.Next, p2, post = p2, post, post.Next
	}
	if p1.Next = post; pre != nil {
		pre.Next, p2 = p2, head
	}
	return p2
}

func main() {}
