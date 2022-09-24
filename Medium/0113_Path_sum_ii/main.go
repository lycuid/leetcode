// https://leetcode.com/problems/path-sum-ii/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ret [][]int) {
	if root != nil {
		if left := pathSum(root.Left, targetSum-root.Val); left != nil {
			for i := range left {
				left[i] = append([]int{root.Val}, left[i]...)
			}
			ret = append(ret, left...)
		}
		if right := pathSum(root.Right, targetSum-root.Val); right != nil {
			for i := range right {
				right[i] = append([]int{root.Val}, right[i]...)
			}
			ret = append(ret, right...)
		}
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			ret = [][]int{{root.Val}}
		}
	}
	return ret
}

func main() {}
