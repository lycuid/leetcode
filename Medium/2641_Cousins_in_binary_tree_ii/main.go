// https://leetcode.com/problems/cousins-in-binary-tree-ii/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelSum(node *TreeNode, lvl int, sum *[]int) {
	if node != nil {
		for len(*sum) <= lvl {
			*sum = append(*sum, 0)
		}
		(*sum)[lvl] += node.Val
		LevelSum(node.Left, lvl+1, sum)
		LevelSum(node.Right, lvl+1, sum)
	}
}

func UpdateTree(parent *TreeNode, lvl int, sum []int) {
	if parent != nil {
		var s int
		if parent.Left != nil {
			s += parent.Left.Val
		}
		if parent.Right != nil {
			s += parent.Right.Val
		}
		if parent.Left != nil {
			parent.Left.Val = sum[lvl] - s
		}
		if parent.Right != nil {
			parent.Right.Val = sum[lvl] - s
		}
		UpdateTree(parent.Left, lvl+1, sum)
		UpdateTree(parent.Right, lvl+1, sum)
	}
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	var sum []int
	LevelSum(root, 0, &sum)
	UpdateTree(root, 1, sum)
	root.Val = 0
	return root
}

func main() {}
