// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(num int, nums ...int) int {
	for i := range nums {
		if nums[i] < num {
			num = nums[i]
		}
	}
	return num
}

type Solution struct{ previous, result int }

func (this *Solution) Aux(root *TreeNode) {
	if root != nil {
		this.Aux(root.Left)
		this.result = Min(this.result, Abs(this.previous-root.Val))
		this.previous = root.Val
		this.Aux(root.Right)
	}
}

func getMinimumDifference(root *TreeNode) int {
	solution := Solution{math.MaxInt, math.MaxInt}
	solution.Aux(root)
	return solution.result
}

func main() {}
