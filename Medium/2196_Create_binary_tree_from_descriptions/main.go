// https://leetcode.com/problems/create-binary-tree-from-descriptions/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) (head *TreeNode) {
	nodes := make(map[int]*TreeNode)
	parentOf := make(map[int]*TreeNode)
	for _, desc := range descriptions {
		parent, child, isLeft := desc[0], desc[1], desc[2] == 1
		if _, found := nodes[parent]; !found {
			nodes[parent] = &TreeNode{Val: parent}
		}
		if _, found := nodes[child]; !found {
			nodes[child] = &TreeNode{Val: child}
		}
		head = nodes[child]
		if isLeft {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
		parentOf[child] = nodes[parent]
	}
	for parentOf[head.Val] != nil {
		head = parentOf[head.Val]
	}
	return head
}

func main() {}
