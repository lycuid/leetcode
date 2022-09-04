// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tuple struct {
	Val, Index, Depth int
}

func Dfs(root *TreeNode, index, depth int, cache *[]Tuple) {
	if root != nil {
		*cache = append(*cache, Tuple{root.Val, index, depth})
		Dfs(root.Left, index-1, depth+1, cache)
		Dfs(root.Right, index+1, depth+1, cache)
	}
}

func verticalTraversal(root *TreeNode) (ret [][]int) {
	var cache []Tuple
	Dfs(root, 0, 0, &cache)
	sort.Slice(cache, func(i, j int) bool {
		if cache[i].Index == cache[j].Index {
			if cache[i].Depth == cache[j].Depth {
				return cache[i].Val < cache[j].Val
			}
			return cache[i].Depth < cache[j].Depth
		}
		return cache[i].Index < cache[j].Index
	})
	Index, tmp := cache[0].Index, []int{}
	for _, tuple := range cache {
		if tuple.Index == Index {
			tmp = append(tmp, tuple.Val)
		} else {
			Index, tmp, ret = tuple.Index, []int{tuple.Val}, append(ret, tmp)
		}
	}
	if len(tmp) > 0 {
		ret = append(ret, tmp)
	}
	return ret
}

func main() {}
