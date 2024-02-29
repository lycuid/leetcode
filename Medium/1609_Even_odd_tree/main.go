// https://leetcode.com/problems/even-odd-tree/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, curr int, levels *[]int) bool {
	if lvl_mod := curr % 2; root != nil {
		if lvl_mod == root.Val%2 {
			return false
		}

		for curr >= len(*levels) {
			val := 0
			if n := len(*levels); n%2 == 1 {
				val = math.MaxInt
			}
			*levels = append(*levels, val)
		}

		if (lvl_mod == 0 && root.Val <= (*levels)[curr]) ||
			(lvl_mod == 1 && root.Val >= (*levels)[curr]) {
			return false
		}
		(*levels)[curr] = root.Val

		return Aux(root.Left, curr+1, levels) && Aux(root.Right, curr+1, levels)
	}
	return true
}

func isEvenOddTree(root *TreeNode) bool {
	levels := []int{}
	return Aux(root, 0, &levels)
}

func main() {}
