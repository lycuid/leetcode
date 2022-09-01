// https://leetcode.com/problems/count-good-nodes-in-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, val int) (ret int) {
	if root == nil {
		return 0
	}
	if root.Val >= val {
		val, ret = root.Val, 1
	}
	return ret + Aux(root.Left, val) + Aux(root.Right, val)
}

func goodNodes(root *TreeNode) int { return Aux(root, root.Val) }

func main() {}
