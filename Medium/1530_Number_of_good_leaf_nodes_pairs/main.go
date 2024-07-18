// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Leaves(root *TreeNode, parent map[*TreeNode]*TreeNode) (ret []*TreeNode) {
	if root != nil {
		if root.Left != nil {
			parent[root.Left] = root
		}
		if root.Right != nil {
			parent[root.Right] = root
		}
		if root.Left == nil && root.Right == nil {
			return append(ret, root)
		}
		ret = append(ret, Leaves(root.Left, parent)...)
		ret = append(ret, Leaves(root.Right, parent)...)
	}
	return ret
}

type Solution struct {
	parent map[*TreeNode]*TreeNode
	leaves map[*TreeNode]bool
}

func MakeSolution(root *TreeNode) Solution {
	parent, leaves := make(map[*TreeNode]*TreeNode), make(map[*TreeNode]bool)
	for _, leaf := range Leaves(root, parent) {
		leaves[leaf] = true
	}
	return Solution{parent, leaves}
}

func (this Solution) Solve(prev, curr *TreeNode, dist int) (ret int) {
	if dist--; dist >= 0 && curr != nil {
		if prev != nil && this.leaves[curr] {
			return 1
		}
		if curr.Left != prev {
			ret += this.Solve(curr, curr.Left, dist)
		}
		if curr.Right != prev {
			ret += this.Solve(curr, curr.Right, dist)
		}
		if this.parent[curr] != prev {
			ret += this.Solve(curr, this.parent[curr], dist)
		}
	}
	return ret
}

func countPairs(root *TreeNode, distance int) (ret int) {
	solution := MakeSolution(root)
	for leaf := range solution.leaves {
		ret += solution.Solve(nil, leaf, distance+1)
		solution.leaves[leaf] = false
	}
	return ret
}

func main() {}
