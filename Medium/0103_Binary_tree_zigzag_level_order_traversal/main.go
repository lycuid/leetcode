// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode, depth int, ret *[][]int) {
	if root != nil {
		if len(*ret) < depth+1 {
			*ret = append(*ret, []int{})
		}
		(*ret)[depth] = append((*ret)[depth], root.Val)
		Solve(root.Left, depth+1, ret)
		Solve(root.Right, depth+1, ret)
	}
}

func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	Solve(root, 0, &ret)
	for i := 1; i < len(ret); i += 2 {
		n := len(ret[i])
		for j := 0; j < n/2; j++ {
			ret[i][j], ret[i][n-j-1] = ret[i][n-j-1], ret[i][j]
		}
	}
	return ret
}

func main() {}
