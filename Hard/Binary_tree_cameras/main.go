// https://leetcode.com/problems/binary-tree-cameras/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode, camera_cnt *int) (self int) {
	if root == nil {
		return 1
	}
	left, right := Solve(root.Left, camera_cnt), Solve(root.Right, camera_cnt)
	if left == 0 || right == 0 {
		*camera_cnt++
		return 2
	}
	if left == 2 || right == 2 {
		return 1
	}
	return 0
}

func minCameraCover(root *TreeNode) (camera_cnt int) {
	if Solve(root, &camera_cnt) == 0 {
		camera_cnt++
	}
	return camera_cnt
}

func main() {}
