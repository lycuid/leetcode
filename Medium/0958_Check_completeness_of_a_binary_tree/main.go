// https://leetcode.com/problems/check-completeness-of-a-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		if root, queue = queue[0], queue[1:]; root == nil {
			for len(queue) > 0 && queue[0] == nil {
				queue = queue[1:]
			}
			break
		}
		queue = append(queue, root.Left)
		queue = append(queue, root.Right)
	}
	return len(queue) == 0
}

func main() {}
