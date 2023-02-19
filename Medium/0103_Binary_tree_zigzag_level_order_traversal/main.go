// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, depth int, ret *[][]int) {
	if root != nil {
		for depth > len(*ret)-1 {
			*ret = append(*ret, []int{})
		}
		if depth%2 == 0 {
			(*ret)[depth] = append((*ret)[depth], root.Val)
		} else {
			(*ret)[depth] = append([]int{root.Val}, (*ret)[depth]...)
		}
		Aux(root.Left, depth+1, ret)
		Aux(root.Right, depth+1, ret)
	}
}

func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	Aux(root, 0, &ret)
	return ret
}

func main() {}
