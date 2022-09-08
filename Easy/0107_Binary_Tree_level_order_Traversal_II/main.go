// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	xs := [][]*TreeNode{{root}}
	for i := 0; i < len(xs); i++ {
		ys := []*TreeNode{}
		for _, val := range xs[i] {
			if val.Left != nil {
				ys = append(ys, val.Left)
			}
			if val.Right != nil {
				ys = append(ys, val.Right)
			}
		}
		if len(ys) > 0 {
			xs = append(xs, ys)
		}
	}

	as := [][]int{}
	for i := len(xs) - 1; i >= 0; i-- {
		bs := []int{}
		for _, val := range xs[i] {
			bs = append(bs, val.Val)
		}
		as = append(as, bs)
	}

	return as
}

func main() {}
