// https://leetcode.com/problems/find-all-groups-of-farmland/
package main

func Aux(grid [][]int, i, j int) []int {
	ret := []int{i, j}
	for ret[0]+1 < len(grid) && grid[ret[0]+1][ret[1]] == 1 {
		ret[0]++
	}
	for ret[1]+1 < len(grid[ret[0]]) && grid[ret[0]][ret[1]+1] == 1 {
		ret[1]++
	}
	for k := j; i <= ret[0]; i++ {
		for j = k; j <= ret[1]; j++ {
			grid[i][j] = 0
		}
	}
	return ret
}

func findFarmland(land [][]int) (ret [][]int) {
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				ret = append(ret, append([]int{i, j}, Aux(land, i, j)...))
			}
		}
	}
	return ret
}

func main() {}
