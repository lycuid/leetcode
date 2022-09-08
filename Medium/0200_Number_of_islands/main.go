// https://leetcode.com/problems/number-of-islands/
package main

const ISLAND = '1'

func Dfs(y, x, id int, grid [][]byte, visited [][]bool) {
	grid[y][x], visited[y][x] = byte(id), true
	if y > 0 && grid[y-1][x] == ISLAND && !visited[y-1][x] {
		Dfs(y-1, x, id, grid, visited)
	}
	if x > 0 && grid[y][x-1] == ISLAND && !visited[y][x-1] {
		Dfs(y, x-1, id, grid, visited)
	}
	if y < len(grid)-1 && grid[y+1][x] == ISLAND && !visited[y+1][x] {
		Dfs(y+1, x, id, grid, visited)
	}
	if x < len(grid[y])-1 && grid[y][x+1] == ISLAND && !visited[y][x+1] {
		Dfs(y, x+1, id, grid, visited)
	}
}

func numIslands(grid [][]byte) (id int) {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == ISLAND && !visited[y][x] {
				id--
				Dfs(y, x, id, grid, visited)
			}
		}
	}
	return -id
}

func main() {}
