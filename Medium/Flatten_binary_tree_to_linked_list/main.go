// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		flatten(root.Left)
		flatten(root.Right)
		for node := root.Left; node != nil; node = node.Right {
			if node.Right == nil {
				node.Right, root.Right, root.Left = root.Right, root.Left, nil
				break
			}
		}
	}
}

func main() {}
