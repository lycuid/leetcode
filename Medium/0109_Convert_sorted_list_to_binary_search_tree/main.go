// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(nodes []int) (root *TreeNode) {
	if n := len(nodes); n > 0 {
		root = &TreeNode{
			Val:   nodes[n/2],
			Left:  Aux(nodes[:n/2]),
			Right: Aux(nodes[n/2+1:]),
		}
	}
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	var nodes []int
	for ; head != nil; head = head.Next {
		nodes = append(nodes, head.Val)
	}
	return Aux(nodes)
}

func main() {}
