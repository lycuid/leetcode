// https://leetcode.com/problems/create-binary-tree-from-descriptions/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) (root *TreeNode) {
	cache := make(map[int]*TreeNode)
	adj := make(map[*TreeNode]*TreeNode)
	for _, d := range descriptions {
		pVal, cVal, isLeft := d[0], d[1], d[2] == 1
		if _, found := cache[pVal]; !found {
			cache[pVal] = &TreeNode{Val: pVal}
		}
		if _, found := cache[cVal]; !found {
			cache[cVal] = &TreeNode{Val: cVal}
		}
		if isLeft {
			cache[pVal].Left = cache[cVal]
		} else {
			cache[pVal].Right = cache[cVal]
		}
		adj[cache[cVal]], root = cache[pVal], cache[cVal]
	}
	for adj[root] != nil {
		root = adj[root]
	}
	return root
}

func main() {}
