// https://leetcode.com/problems/trim-a-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getRoot(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		goto EXIT
	}
	if root.Val < low {
		return getRoot(root.Right, low, high)
	}
	if root.Val > high {
		return getRoot(root.Left, low, high)
	}
EXIT:
	return root
}

func trimBST(t *TreeNode, low, high int) *TreeNode {
	root := getRoot(t, low, high)
	if root == nil {
		return nil
	}
	l, r := root, root
	for l != nil && l.Left != nil {
		if l.Left.Val < low {
			l.Left = l.Left.Right
			continue
		}
		l = l.Left
	}
	if l != nil && l.Left != nil {
		l.Left.Left = nil
	}
	for r != nil && r.Right != nil {
		if r.Right.Val > high {
			r.Right = r.Right.Left
			continue
		}
		r = r.Right
	}
	if r != nil && r.Right != nil {
		r.Right.Right = nil
	}
	return root
}

func main() {}
