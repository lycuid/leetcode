// https://leetcode.com/problems/binary-tree-level-order-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode, depth int, ret *[][]int) {
	if root != nil {
		if depth >= len(*ret) {
			*ret = append(*ret, []int{})
		}
		(*ret)[depth] = append((*ret)[depth], root.Val)
		Solve(root.Left, depth+1, ret)
		Solve(root.Right, depth+1, ret)
	}
}

func levelOrder(root *TreeNode) (ret [][]int) {
	Solve(root, 0, &ret)
	return ret
}

func main() {}
