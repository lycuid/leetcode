// https://leetcode.com/problems/range-sum-of-bst/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(t *TreeNode, l int, h int) (n int) {
	if t != nil {
		if n = rangeSumBST(t.Left, l, h) + rangeSumBST(t.Right, l, h); t.Val >= l && t.Val <= h {
			n += t.Val
		}
	}
	return n
}

func main() {}
