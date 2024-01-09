// https://leetcode.com/problems/leaf-similar-trees/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CollectLeaves(root *TreeNode, leaves *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*leaves = append(*leaves, root.Val)
		}
		CollectLeaves(root.Left, leaves)
		CollectLeaves(root.Right, leaves)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) (ret bool) {
	var leaves1, leaves2 []int
	CollectLeaves(root1, &leaves1)
	CollectLeaves(root2, &leaves2)
	for i := 0; i < len(leaves1) && i < len(leaves2); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return len(leaves1) == len(leaves2)
}

func main() {}
