// https://leetcode.com/problems/evaluate-boolean-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) (ret bool) {
	if root != nil {
		switch root.Val {
		case 0:
			return false
		case 1:
			return true
		case 2:
			return evaluateTree(root.Left) || evaluateTree(root.Right)
		case 3:
			return evaluateTree(root.Left) && evaluateTree(root.Right)
		}
	}
	return ret
}

func main() {}
