// https://leetcode.com/problems/find-bottom-left-tree-value/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode) (leaf int, depth int) {
	if root != nil {
		leaf, depth = Aux(root.Left)
		if rl, rd := Aux(root.Right); rd > depth {
			leaf, depth = rl, rd
		}
		if depth++; depth == 1 {
			leaf = root.Val
		}
	}
	return leaf, depth
}

func findBottomLeftValue(root *TreeNode) int {
	leaf, _ := Aux(root)
	return leaf
}

func main() {}
