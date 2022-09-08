// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
package main

type Node struct {
	Val      int
	Children []*Node
}

func Dfs(root *Node, depth int, values *[][]int) {
	if root != nil {
		for len(*values) <= depth {
			*values = append(*values, []int{})
		}
		(*values)[depth] = append((*values)[depth], root.Val)
		for _, child := range root.Children {
			Dfs(child, depth+1, values)
		}
	}
}

func levelOrder(root *Node) (ret [][]int) {
	Dfs(root, 0, &ret)
	return ret
}

func main() {}
