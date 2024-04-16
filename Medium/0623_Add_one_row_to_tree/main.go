// https://leetcode.com/problems/add-one-row-to-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root != nil {
		switch depth {
		case 1:
			return &TreeNode{Val: val, Left: root}
		case 2:
			root.Left = &TreeNode{Val: val, Left: root.Left}
			root.Right = &TreeNode{Val: val, Right: root.Right}
		default:
			root.Left = addOneRow(root.Left, val, depth-1)
			root.Right = addOneRow(root.Right, val, depth-1)
		}
	}
	return root
}

func main() {}
