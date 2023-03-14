// https://leetcode.com/problems/sum-root-to-leaf-numbers/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Aux(root *TreeNode, n int) int {
	if root != nil {
		n = n*10 + root.Val
		return Max(n, Aux(root.Left, n)+Aux(root.Right, n))
	}
	return 0
}

func sumNumbers(root *TreeNode) int { return Aux(root, 0) }

func main() {}
