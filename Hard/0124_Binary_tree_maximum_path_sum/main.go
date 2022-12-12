// https://leetcode.com/problems/binary-tree-maximum-path-sum/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(nums ...int) int {
	num := nums[0]
	for i := range nums {
		if nums[i] > num {
			num = nums[i]
		}
	}
	return num
}

func Aux(root *TreeNode) (int, int) {
	if root != nil {
		aleft, dleft := Aux(root.Left)
		aright, dright := Aux(root.Right)
		return Max(root.Val, root.Val+Max(0, aleft), root.Val+Max(0, aright)),
			Max(dleft, dright, root.Val+Max(0, aleft)+Max(0, aright))
	}
	return math.MinInt, math.MinInt
}

func maxPathSum(root *TreeNode) int { return Max(Aux(root)) }

func main() {}
