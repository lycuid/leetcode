// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
package main

type Item struct {
	y, x, walls int
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	steps, done := make(map[Item]int), make(map[Item]bool)
	for q := []Item{{0, 0, k}}; len(q) > 0; q = q[1:] {
		item, walls := q[0], q[0].walls
		if done[item] || item.walls < 0 {
			continue
		}
		if item.y == m-1 && item.x == n-1 {
			return steps[item]
		}
		if done[item] = true; grid[item.y][item.x] == 1 {
			walls--
		}
		for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			dy, dx := item.y+d[0], item.x+d[1]
			if dy >= 0 && dy < m && dx >= 0 && dx < n {
				new_item := Item{dy, dx, walls}
				q, steps[new_item] = append(q, new_item), steps[item]+1
			}
		}
	}
	return -1
}

func main() {}
