// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Accumulate(root *TreeNode) int {
	if root != nil {
		root.Val += Accumulate(root.Left) + Accumulate(root.Right)
		return root.Val
	}
	return 0
}

func Solve(root *TreeNode, val int) int {
	if root != nil {
		return max(
			((val - root.Val) * root.Val),
			Solve(root.Left, val),
			Solve(root.Right, val),
		)
	}
	return 0
}

func maxProduct(root *TreeNode) int {
	return Solve(root, Accumulate(root)) % (1e9 + 7)
}

func main() {}
