// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
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

func Sum(root *TreeNode, sum_at map[*TreeNode]int) int {
	if root != nil {
		sum_at[root] = root.Val + Sum(root.Left, sum_at) + Sum(root.Right, sum_at)
	}
	return sum_at[root]
}

func Aux(root *TreeNode, sum_at map[*TreeNode]int, sum int) (ret int) {
	if root != nil {
		ret = Max(
			sum_at[root]*(sum-sum_at[root]),
			Max(Aux(root.Left, sum_at, sum), Aux(root.Right, sum_at, sum)),
		)
	}
	return ret
}

func maxProduct(root *TreeNode) int {
	sum_at := make(map[*TreeNode]int)
	return Aux(root, sum_at, Sum(root, sum_at)) % (1e9 + 7)
}

func main() {}
