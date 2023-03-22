// https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/
package main

import "math"

func Aux(src, dst int, edges map[int]map[int]int, visited map[int]bool) (ret int) {
	if ret = math.MaxInt; !visited[src] {
		visited[src] = true
		for child, distance := range edges[src] {
			if distance < ret {
				ret = distance
			}
			if n := Aux(child, dst, edges, visited); n < ret {
				ret = n
			}
		}
	}
	return ret
}

func minScore(n int, roads [][]int) int {
	edges := make(map[int]map[int]int)
	for _, e := range roads {
		from, to, distance := e[0], e[1], e[2]
		if edges[from] == nil {
			edges[from] = make(map[int]int)
		}
		if edges[to] == nil {
			edges[to] = make(map[int]int)
		}
		edges[from][to], edges[to][from] = distance, distance
	}
	return Aux(1, n, edges, make(map[int]bool))
}

func main() {}
