// https://leetcode.com/problems/paint-house-iii/
package main

const INF = 1e9 + 7

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCost(houses []int, cost [][]int, m int, n int, target int) (ret int) {
	cache := make([][]int, target+1)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = INF
		}
	}
	if houses[0] == 0 {
		cache[1] = cost[0]
	} else {
		cache[1][houses[0]-1] = 0
	}
	for h := 1; h < m; h++ {
		for group := target; group > 0; group-- {
			for color := 0; color < n; color++ {
				min_cost := cache[group][color]
				cache[group][color] = INF
				if houses[h] != 0 && houses[h] != color+1 {
					continue
				}
				c := 0
				if houses[h] == 0 {
					c = cost[h][color]
				}
				min_cost += c
				for color1 := 0; color1 < n; color1++ {
					if color == color1 {
						continue
					}
					min_cost = Min(min_cost, cache[group-1][color1]+c)
				}
				cache[group][color] = min_cost
			}
		}
	}
	ret = INF
	for i := 0; i < n; i++ {
		ret = Min(ret, cache[target][i])
	}
	if ret < INF {
		return ret
	}
	return -1
}

func main() {}
