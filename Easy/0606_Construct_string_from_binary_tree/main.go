// https://leetcode.com/problems/construct-string-from-binary-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(n *TreeNode) (ret string) {
	if n != nil {
		if ret = fmt.Sprintf("%d", n.Val); n.Left != nil || n.Right != nil {
			ret += ("(" + tree2str(n.Left) + ")")
			if n.Right != nil {
				ret += ("(" + tree2str(n.Right) + ")")
			}
		}
	}
	return ret
}

func main() {}
