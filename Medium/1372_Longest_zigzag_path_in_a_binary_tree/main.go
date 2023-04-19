// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(num int, nums ...int) int {
	for i := range nums {
		if nums[i] > num {
			num = nums[i]
		}
	}
	return num
}

func Aux(root *TreeNode, from_left bool) (int, int) {
	if root == nil {
		return -1, -1
	}
	lcont, laux := Aux(root.Left, false)
	rcont, raux := Aux(root.Right, true)
	if from_left {
		return lcont + 1, Max(laux, rcont+1, raux)
	}
	return rcont + 1, Max(raux, lcont+1, laux)
}

func longestZigZag(root *TreeNode) int {
	return Max(Max(Aux(root, false)), Max(Aux(root, true)))
}

func main() {}
