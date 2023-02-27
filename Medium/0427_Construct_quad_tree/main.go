// https://leetcode.com/problems/construct-quad-tree/
package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func Aux(y, x, n int, grid [][]int) *Node {
	var node Node
	if n > 0 {
		s := grid[y+n-1][x+n-1] - grid[y+n-1][x-1] - grid[y-1][x+n-1] + grid[y-1][x-1]
		if node.IsLeaf, node.Val = s == 0 || s == (n*n), s != 0; !node.IsLeaf {
			m := n >> 1
			node.TopLeft = Aux(y, x, m, grid)
			node.TopRight = Aux(y, x+m, m, grid)
			node.BottomLeft = Aux(y+m, x, m, grid)
			node.BottomRight = Aux(y+m, x+m, m, grid)
		}
	}
	return &node
}

func construct(grid [][]int) *Node {
	sum, n := make([][]int, len(grid)+1), len(grid)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = grid[i-1][j-1] + sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1]
		}
	}
	return Aux(1, 1, n, sum)
}

func main() {}
