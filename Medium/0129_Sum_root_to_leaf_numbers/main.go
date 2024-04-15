// https://leetcode.com/problems/sum-root-to-leaf-numbers/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) (sum int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		if root.Left != nil {
			root.Left.Val += root.Val * 10
			sum += sumNumbers(root.Left)
		}
		if root.Right != nil {
			root.Right.Val += root.Val * 10
			sum += sumNumbers(root.Right)
		}
	}
	return sum
}

func main() {}
