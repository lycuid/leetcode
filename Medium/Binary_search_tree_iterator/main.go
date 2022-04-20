// https://leetcode.com/problems/binary-search-tree-iterator/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
}

func inorderTree(root *TreeNode) *TreeNode {
	r, old_root := root, root
	if root == nil {
		goto EXIT
	}
	root.Right = inorderTree(root.Right)
	if root.Left == nil {
		goto EXIT
	}
	root = inorderTree(root.Left)
	old_root.Left = nil
	for r = root; r.Right != nil; r = r.Right {
	}
	r.Right = old_root
EXIT:
	return root
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: inorderTree(root)}
}

func (this *BSTIterator) Next() (ret int) {
	ret, this.root = this.root.Val, this.root.Right
	return ret
}

func (this *BSTIterator) HasNext() bool {
	return this.root != nil
}

func main() {}
