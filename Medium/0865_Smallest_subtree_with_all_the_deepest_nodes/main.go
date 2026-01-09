// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Depth(root *TreeNode, cache map[*TreeNode]int) int {
	if root != nil {
		cache[root] = max(Depth(root.Left, cache), Depth(root.Right, cache)) + 1
	}
	return cache[root]
}

func Solve(root *TreeNode, cache map[*TreeNode]int) *TreeNode {
	if left, right := cache[root.Left], cache[root.Right]; left > right {
		return Solve(root.Left, cache)
	} else if left < right {
		return Solve(root.Right, cache)
	}
	return root
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	cache := make(map[*TreeNode]int)
	Depth(root, cache)
	return Solve(root, cache)
}

func main() {}
