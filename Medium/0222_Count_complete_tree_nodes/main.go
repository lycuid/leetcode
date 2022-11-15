// https://leetcode.com/problems/count-complete-tree-nodes/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) (ret int) {
	if root == nil {
		return 0
	}
	depthLeft, depthRight := 0, 0
	for head := root; head != nil; head = head.Left {
		depthLeft++
	}
	for head := root; head != nil; head = head.Right {
		depthRight++
	}
	if depthLeft == depthRight {
		return (1 << depthLeft) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {}
