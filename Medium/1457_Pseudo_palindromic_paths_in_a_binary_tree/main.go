// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
package main

import "math/bits"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, mask uint16) int {
	if root == nil {
		return 0
	}
	if mask ^= 1 << root.Val; root.Left == nil && root.Right == nil {
		if bits.OnesCount16(mask) > 1 {
			return 0
		}
		return 1
	}
	return Aux(root.Left, mask) + Aux(root.Right, mask)
}

func pseudoPalindromicPaths(root *TreeNode) int {
	return Aux(root, 0)
}

func main() {}
