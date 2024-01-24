// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
package main

import "math/bits"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, mask uint16) (ret int) {
	if root == nil {
		return 0
	}
	mask ^= 1 << root.Val
	if root.Left == nil && root.Right == nil && bits.OnesCount16(mask) <= 1 {
		return 1
	}
	return Aux(root.Left, mask) + Aux(root.Right, mask)
}

func pseudoPalindromicPaths(root *TreeNode) (ret int) {
	return Aux(root, 0)
}

func main() {}
