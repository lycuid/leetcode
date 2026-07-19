// https://leetcode.com/contest/weekly-contest-511/problems/count-dominant-nodes-in-a-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countDominantNodes(root *TreeNode) (count int) {
	if root != nil {
		count += countDominantNodes(root.Left)
		count += countDominantNodes(root.Right)
		var mx int
		if root.Left != nil {
			mx = max(mx, root.Left.Val)
		}
		if root.Right != nil {
			mx = max(mx, root.Right.Val)
		}
		if root.Val >= mx {
			count++
		} else {
			root.Val = mx
		}
	}
	return count
}

func main() {}
