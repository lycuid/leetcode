// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
package main

func countUnguarded(m int, n int, guards [][]int, walls [][]int) (count int) {
	cache := make([][]uint8, m)
	for i := range cache {
		cache[i] = make([]uint8, n)
	}
	for _, w := range walls {
		cache[w[0]][w[1]] = 2
	}
	for _, g := range guards {
		cache[g[0]][g[1]] |= 1
		for x := g[0] + 1; x < m && cache[x][g[1]]&0b001011 == 0; x++ {
			cache[x][g[1]] |= 0b001000
		}
		for x := g[0] - 1; x >= 0 && cache[x][g[1]]&0b100011 == 0; x-- {
			cache[x][g[1]] |= 0b100000
		}
		for y := g[1] + 1; y < n && cache[g[0]][y]&0b010011 == 0; y++ {
			cache[g[0]][y] |= 0b010000
		}
		for y := g[1] - 1; y >= 0 && cache[g[0]][y]&0b000111 == 0; y-- {
			cache[g[0]][y] |= 0b000100
		}
	}
	for _, row := range cache {
		for _, col := range row {
			if col == 0 {
				count++
			}
		}
	}
	return count
}

func main() {}
