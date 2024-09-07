// https://leetcode.com/problems/linked-list-in-binary-tree/
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

func Aux(head *ListNode, node *TreeNode) bool {
	if head == nil {
		return true
	}
	if node == nil {
		return false
	}
	if head.Val != node.Val {
		return false
	}
	return Aux(head.Next, node.Left) || Aux(head.Next, node.Right)
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		if node := stack[0]; node != nil {
			if Aux(head, node) {
				return true
			}
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
		stack = stack[1:]
	}
	return false
}

func main() {}
