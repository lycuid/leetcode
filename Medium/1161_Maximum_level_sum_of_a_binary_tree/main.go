// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct{ values []int }

func (this *Solution) Aux(root *TreeNode, level int) {
	if root != nil {
		for len(this.values) <= level {
			this.values = append(this.values, 0)
		}
		this.values[level] += root.Val
		this.Aux(root.Left, level+1)
		this.Aux(root.Right, level+1)
	}
}

func (this Solution) MaxIndex() int {
	ret, index := this.values[1], 1
	for i := 2; i < len(this.values); i++ {
		if this.values[i] > ret {
			ret, index = this.values[i], i
		}
	}
	return index
}

func maxLevelSum(root *TreeNode) int {
	solution := Solution{}
	solution.Aux(root, 1)
	return solution.MaxIndex()
}

func main() {}
