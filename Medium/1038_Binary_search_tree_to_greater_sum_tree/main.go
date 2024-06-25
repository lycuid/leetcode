// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	var Aux func(*TreeNode, int) int
	Aux = func(root *TreeNode, baseVal int) (sum int) {
		if root != nil {
			root.Val += Aux(root.Right, baseVal)
			sum = root.Val + Aux(root.Left, root.Val+baseVal)
			root.Val += baseVal
		}
		return sum
	}
	Aux(root, 0)
	return root
}

func main() {}
