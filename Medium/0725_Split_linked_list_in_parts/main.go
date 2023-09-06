// https://leetcode.com/problems/split-linked-list-in-parts/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	ret, n := make([]*ListNode, k), 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	for i := range ret {
		count := n / k
		if i < n%k {
			count++
		}
		node := head
		for ret[i] = head; count > 1; count-- {
			node = node.Next
		}
		if node != nil {
			head, node.Next = node.Next, nil
		}
	}
	return ret
}

func main() {}
