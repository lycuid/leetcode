// https://leetcode.com/problems/balance-a-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	var (
		nodes     []*TreeNode
		Inorder   func(*TreeNode) []*TreeNode
		BuildTree func([]*TreeNode) *TreeNode
	)

	Inorder = func(root *TreeNode) []*TreeNode {
		if root != nil {
			Inorder(root.Left)
			nodes = append(nodes, root)
			Inorder(root.Right)
		}
		return nodes
	}

	BuildTree = func(nodes []*TreeNode) (node *TreeNode) {
		if n := len(nodes); n > 0 {
			node = nodes[n/2]
			node.Left = BuildTree(nodes[:n/2])
			node.Right = BuildTree(nodes[n/2+1:])
		}
		return node
	}

	return BuildTree(Inorder(root))
}

func main() {}
