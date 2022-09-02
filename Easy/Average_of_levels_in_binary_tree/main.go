// https://leetcode.com/problems/average-of-levels-in-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (ret []float64) {
	queue := []*TreeNode{root}
	for level, cursor := 0, 1; cursor > 0; level, cursor = level+1, len(queue) {
		ret = append(ret, 0)
		for q := 0; q < cursor; q++ {
			ret[level] += float64(queue[q].Val)
			if queue[q].Left != nil {
				queue = append(queue, queue[q].Left)
			}
			if queue[q].Right != nil {
				queue = append(queue, queue[q].Right)
			}
		}
		ret[level], queue = ret[level]/float64(cursor), queue[cursor:]
	}
	return ret
}

func main() {}
