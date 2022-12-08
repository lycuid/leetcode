// https://leetcode.com/problems/leaf-similar-trees/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Leaves(root *TreeNode, leaves *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*leaves = append(*leaves, root.Val)
		}
		Leaves(root.Left, leaves)
		Leaves(root.Right, leaves)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var l1, l2 []int
	Leaves(root1, &l1)
	Leaves(root2, &l2)
	for len(l1) > 0 && len(l2) > 0 && l1[0] == l2[0] {
		l1, l2 = l1[1:], l2[1:]
	}
	return len(l1)+len(l2) == 0
}

func main() {}
