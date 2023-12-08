// https://leetcode.com/problems/construct-string-from-binary-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) (str string) {
	if root != nil {
		str = fmt.Sprintf("%d", root.Val)
		if root.Left != nil || root.Right != nil {
			str += fmt.Sprintf("(%s)", tree2str(root.Left))
		}
		if root.Right != nil {
			str += fmt.Sprintf("(%s)", tree2str(root.Right))
		}
	}
	return str
}

func main() {}
