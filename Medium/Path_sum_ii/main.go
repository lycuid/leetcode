// https://leetcode.com/problems/path-sum-ii/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, t int) (ret [][]int) {
	if root == nil {
		return nil
	}
	if root.Val == t && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	left := pathSum(root.Left, t-root.Val)
	if left != nil {
		for i := 0; i < len(left); i++ {
			left[i] = append([]int{root.Val}, left[i]...)
		}
		ret = append(ret, left...)
	}
	right := pathSum(root.Right, t-root.Val)
	if right != nil {
		for i := 0; i < len(right); i++ {
			right[i] = append([]int{root.Val}, right[i]...)
		}
		ret = append(ret, right...)
	}
	return ret
}

func main() {}
