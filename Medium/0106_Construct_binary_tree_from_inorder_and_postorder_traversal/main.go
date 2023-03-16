// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) (head *TreeNode) {
	for count := 0; len(inorder) > 0; count = 0 {
		for postorder[count] != inorder[0] {
			count++
		}
		head = &TreeNode{
			Val:   inorder[0],
			Left:  head,
			Right: buildTree(inorder[1:count+1], postorder[:count]),
		}
		inorder, postorder = inorder[count+1:], postorder[count+1:]
	}
	return head
}

func main() {}
