// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, k int, seen map[int]bool) bool {
	if root == nil {
		return false
	}
	if seen[k-root.Val] {
		return true
	}
	seen[root.Val] = true
	return Aux(root.Left, k, seen) || Aux(root.Right, k, seen)
}

func findTarget(root *TreeNode, k int) bool {
	return Aux(root, k, make(map[int]bool))
}

func main() {}
