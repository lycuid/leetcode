// https://leetcode.com/problems/range-sum-of-bst/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) (sum int) {
	if root != nil {
		sum = rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
	}
	return sum
}

func main() {}
