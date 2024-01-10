// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeGraph(root *TreeNode, adj *map[int][]int) {
	if root != nil {
		if root.Left != nil {
			(*adj)[root.Val] = append((*adj)[root.Val], root.Left.Val)
			(*adj)[root.Left.Val] = append((*adj)[root.Left.Val], root.Val)
			MakeGraph(root.Left, adj)
		}
		if root.Right != nil {
			(*adj)[root.Val] = append((*adj)[root.Val], root.Right.Val)
			(*adj)[root.Right.Val] = append((*adj)[root.Right.Val], root.Val)
			MakeGraph(root.Right, adj)
		}
	}
}

func Depth(parent int, adj map[int][]int, visited map[int]bool) (depth int) {
	visited[parent] = true
	for _, child := range adj[parent] {
		if !visited[child] {
			if d := 1 + Depth(child, adj, visited); d > depth {
				depth = d
			}
		}
	}
	return depth
}

func amountOfTime(root *TreeNode, start int) int {
	adj, visited := make(map[int][]int), make(map[int]bool)
	MakeGraph(root, &adj)
	return Depth(start, adj, visited)
}

func main() {}
