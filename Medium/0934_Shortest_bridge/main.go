// https://leetcode.com/problems/shortest-bridge/
package main

import "math"

type Tuple struct{ y, x int }

func Visit(y, x, island_no int, grid [][]int, island *[4][]Tuple) {
	if y >= len(grid) || y < 0 || x >= len(grid[0]) || x < 0 || grid[y][x] != 1 {
		return
	}
	island[island_no], grid[y][x] = append(island[island_no], Tuple{y, x}), 0
	Visit(y+1, x, island_no, grid, island)
	Visit(y-1, x, island_no, grid, island)
	Visit(y, x+1, island_no, grid, island)
	Visit(y, x-1, island_no, grid, island)
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func shortestBridge(grid [][]int) int {
	var island [4][]Tuple
	island_no, ret := 2, math.MaxInt
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				Visit(i, j, island_no, grid, &island)
				island_no++
			}
		}
	}
	for _, i1 := range island[2] {
		for _, i2 := range island[3] {
			ret = Min(ret, Abs(i1.y-i2.y)+Abs(i1.x-i2.x))
		}
	}
	return ret - 1
}

func main() {}
