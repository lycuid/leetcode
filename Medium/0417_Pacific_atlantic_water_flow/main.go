// https://leetcode.com/problems/pacific-atlantic-water-flow/
package main

const (
	North = 1 << iota
	South = 1 << iota
	East  = 1 << iota
	West  = 1 << iota
)

func Dfs(i, j int, heights, dir [][]int, visited [][]bool) int {
	visited[i][j] = true
	for _, d := range [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
		di, dj := i+d[0], j+d[1]
		if di >= 0 && dj >= 0 && di < len(dir) && dj < len(dir[0]) {
			if heights[di][dj] <= heights[i][j] {
				if !visited[di][dj] {
					Dfs(di, dj, heights, dir, visited)
				}
				if dir[i][j] |= dir[di][dj]; heights[di][dj] == heights[i][j] {
					dir[di][dj] |= dir[i][j]
				}
			}
		}
	}
	return dir[i][j]
}

func pacificAtlantic(heights [][]int) (ret [][]int) {
	m, n := len(heights), len(heights[0])
	dir, visited := make([][]int, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		dir[i], visited[i] = make([]int, n), make([]bool, n)
	}
	for i := 0; i < m; i++ {
		dir[i][0] |= West
		dir[i][n-1] |= East
	}
	for j := 0; j < n; j++ {
		dir[0][j] |= North
		dir[m-1][j] |= South
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			Dfs(i, j, heights, dir, visited)
			if dir[i][j]&(North|West) > 0 && dir[i][j]&(South|East) > 0 {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret
}

func main() {}
