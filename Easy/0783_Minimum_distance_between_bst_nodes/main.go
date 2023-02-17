// https://leetcode.com/problems/minimum-distance-between-bst-nodes/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(n int, nums ...int) int {
	for _, num := range nums {
		if num < n {
			n = num
		}
	}
	return n
}

func Max(n int, nums ...int) int {
	for _, num := range nums {
		if num > n {
			n = num
		}
	}
	return n
}

type Tuple struct{ min, max, diff int }

func Aux(root *TreeNode) (ret Tuple) {
	if ret = (Tuple{math.MaxInt, -1, math.MaxInt}); root != nil {
		l, r := Aux(root.Left), Aux(root.Right)
		if root.Left != nil {
			ret.diff = Min(ret.diff, root.Val-l.max)
		}
		if root.Right != nil {
			ret.diff = Min(ret.diff, r.min-root.Val)
		}
		ret.min = Min(ret.min, l.min, root.Val)
		ret.max = Max(ret.max, r.max, root.Val)
		ret.diff = Min(ret.diff, l.diff, r.diff)
	}
	return ret
}

func minDiffInBST(root *TreeNode) (ret int) { return Aux(root).diff }

func main() {}
