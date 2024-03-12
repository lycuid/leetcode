// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) (root *ListNode) {
	root = &ListNode{Next: head}
	for outter, sum_o := root, 0; outter != nil; outter = outter.Next {
		sum_o += outter.Val
		for inner, sum_i := root, 0; inner != outter; inner = inner.Next {
			if sum_i += inner.Val; sum_i == sum_o {
				inner.Next = outter.Next
				break
			}
		}
	}
	return root.Next
}

func main() {}
