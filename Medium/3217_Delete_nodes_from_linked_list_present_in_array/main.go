// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Aux(node *ListNode, cache map[int]bool) *ListNode {
	if node != nil {
		if node.Next = Aux(node.Next, cache); cache[node.Val] {
			return node.Next
		}
	}
	return node
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	cache := make(map[int]bool)
	for _, num := range nums {
		cache[num] = true
	}
	return Aux(head, cache)
}

func main() {}
