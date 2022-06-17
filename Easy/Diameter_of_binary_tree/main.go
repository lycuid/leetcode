// https://leetcode.com/problems/diameter-of-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Depth(root *TreeNode, depth *map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	(*depth)[root] = 1 + Max(Depth(root.Left, depth), Depth(root.Right, depth))
	return (*depth)[root]
}

func Aux(root *TreeNode, depth *map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	return Max((*depth)[root.Left]+(*depth)[root.Right], Max(Aux(root.Left, depth), Aux(root.Right, depth)))
}

func diameterOfBinaryTree(root *TreeNode) int {
	depth := make(map[*TreeNode]int)
	depth[root] = Depth(root, &depth)
	return Aux(root, &depth)
}

func main() {}
