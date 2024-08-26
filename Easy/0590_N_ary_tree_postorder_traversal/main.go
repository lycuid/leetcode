// https://leetcode.com/problems/n-ary-tree-postorder-traversal/
package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (values []int) {
	if root != nil {
		for _, child := range root.Children {
			values = append(values, postorder(child)...)
		}
		values = append(values, root.Val)
	}
	return values
}

func main() {}
