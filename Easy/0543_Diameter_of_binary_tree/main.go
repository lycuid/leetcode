// https://leetcode.com/problems/diameter-of-binary-tree/
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

func Aux(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lmax, ldome := Aux(root.Left)
	rmax, rdome := Aux(root.Right)
	return 1 + Max(lmax, rmax), Max(ldome, rdome, 1+lmax+rmax)
}

func diameterOfBinaryTree(root *TreeNode) int {
	max, dome := Aux(root)
	return Max(max, dome) - 1
}

func main() {}
