// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) (root *TreeNode) {
	for p := 0; p < len(preorder) && root == nil; p++ {
		for i := 0; i < len(inorder) && root == nil; i++ {
			if preorder[p] == inorder[i] {
				root = &TreeNode{
					Val:   preorder[p],
					Left:  buildTree(preorder[p:], inorder[:i]),
					Right: buildTree(preorder[p:], inorder[i+1:]),
				}
			}
		}
	}
	return root
}

func main() {}
