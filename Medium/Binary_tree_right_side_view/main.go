// https://leetcode.com/problems/binary-tree-right-side-view/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode, depth int, ret *[]int) {
	if root != nil {
		if depth < len(*ret) {
			(*ret)[depth] = root.Val
		} else {
			*ret = append(*ret, root.Val)
		}
		Solve(root.Left, depth+1, ret)
		Solve(root.Right, depth+1, ret)
	}
}

func rightSideView(root *TreeNode) (ret []int) {
	Solve(root, 0, &ret)
	return ret
}

func main() {}
