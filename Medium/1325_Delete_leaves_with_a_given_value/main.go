// https://leetcode.com/problems/delete-leaves-with-a-given-value/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root != nil {
		root.Left = removeLeafNodes(root.Left, target)
		root.Right = removeLeafNodes(root.Right, target)
		if root.Val == target && root.Left == nil && root.Right == nil {
			return nil
		}
	}
	return root
}

func main() {}
