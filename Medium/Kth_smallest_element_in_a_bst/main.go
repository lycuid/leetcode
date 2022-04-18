// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root *TreeNode, k *int) *TreeNode {
	if root == nil {
		return nil
	}
	left := solve(root.Left, k)
	if left != nil {
		return left
	}
	(*k)--
	if (*k) == 0 {
		return root
	}
	return solve(root.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
	new_root := solve(root, &k)
	if new_root != nil {
		return new_root.Val
	}
	return 0
}

func main() {}
