// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(node *TreeNode, lvl int, sum *[]int64) {
	if node != nil {
		for len(*sum) <= lvl {
			*sum = append(*sum, 0)
		}
		(*sum)[lvl] += int64(node.Val)
		Aux(node.Left, lvl+1, sum)
		Aux(node.Right, lvl+1, sum)
	}
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var sum []int64
	if Aux(root, 0, &sum); k > len(sum) {
		return -1
	}
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	return sum[k-1]
}

func main() {}
