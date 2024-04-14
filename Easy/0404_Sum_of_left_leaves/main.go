// https://leetcode.com/problems/sum-of-left-leaves/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (sum int) {
	if root != nil {
		sum = sumOfLeftLeaves(root.Right) + sumOfLeftLeaves(root.Left)
		if node := root.Left; node != nil && node.Left == nil && node.Right == nil {
			sum += node.Val
		}
	}
	return sum
}

func main() {}
