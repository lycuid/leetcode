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
			return &TreeNode{val, root, nil}
		case 2:
			root.Left = &TreeNode{val, root.Left, nil}
			root.Right = &TreeNode{val, nil, root.Right}
		default:
			addOneRow(root.Left, val, depth-1)
			addOneRow(root.Right, val, depth-1)
		}
	}
	return root
}

func main() {}
