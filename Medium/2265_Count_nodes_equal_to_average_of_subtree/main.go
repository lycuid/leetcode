// https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	leftRes, leftSize := solve(root.Left)
	rightRes, rightSize := solve(root.Right)
	size, res := leftSize+rightSize+1, leftRes+rightRes

	val := root.Val
	if root.Left != nil {
		root.Val += root.Left.Val
	}
	if root.Right != nil {
		root.Val += root.Right.Val
	}
	if root.Val/size == val {
		res++
	}
	return res, size
}

func averageOfSubtree(root *TreeNode) int {
	res, _ := solve(root)
	return res
}

func main() {}
