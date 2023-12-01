// https://leetcode.com/problems/find-mode-in-binary-search-tree/
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tuple struct{ Val, Count int }

func inorder(root *TreeNode, items *[]Tuple) {
	if root != nil {
		inorder(root.Left, items)
		if n := len(*items); n > 0 && (*items)[n-1].Val == root.Val {
			(*items)[n-1].Count++
		} else {
			*items = append(*items, Tuple{root.Val, 1})
		}
		inorder(root.Right, items)
	}
}

func findMode(root *TreeNode) (ret []int) {
	var items []Tuple
	inorder(root, &items)
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})
	for i := 0; i < len(items) && items[i].Count == items[0].Count; i++ {
		ret = append(ret, items[i].Val)
	}
	return ret
}

func main() {}
