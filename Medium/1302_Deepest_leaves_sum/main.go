// https://leetcode.com/problems/deepest-leaves-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func DeepestLeaf(root *TreeNode, depth int) int {
	if root == nil {
		return 0
	}
	if depth == 1 {
		return root.Val
	}
	return DeepestLeaf(root.Left, depth-1) + DeepestLeaf(root.Right, depth-1)
}

func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(Depth(root.Left), Depth(root.Right))
}

func deepestLeavesSum(root *TreeNode) int {
	return DeepestLeaf(root, Depth(root))
}

func main() {}
