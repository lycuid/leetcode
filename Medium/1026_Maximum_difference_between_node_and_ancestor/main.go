// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Aux(root *TreeNode, min, max int) (ret int) {
	if root != nil {
		min, max = Min(min, root.Val), Max(max, root.Val)
		ret = Max(Aux(root.Left, min, max), Aux(root.Right, min, max))
	}
	return Max(ret, max-min)
}

func maxAncestorDiff(root *TreeNode) int {
	return Aux(root, math.MaxInt, 0)
}

func main() {}
