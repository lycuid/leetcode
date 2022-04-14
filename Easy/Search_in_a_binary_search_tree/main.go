// https://leetcode.com/problems/search-in-a-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		goto EXIT
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}
EXIT:
	return root
}

func main() {}
