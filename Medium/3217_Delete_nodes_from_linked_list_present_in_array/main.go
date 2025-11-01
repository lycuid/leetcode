// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	var (
		cache  = make(map[int]bool)
		remove func(*ListNode) *ListNode
	)

	for _, num := range nums {
		cache[num] = true
	}
	remove = func(head *ListNode) *ListNode {
		if head != nil {
			if head.Next = remove(head.Next); cache[head.Val] {
				return head.Next
			}
		}
		return head
	}

	return remove(head)
}

func main() {}
