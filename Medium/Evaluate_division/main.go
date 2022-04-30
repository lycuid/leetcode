// https://leetcode.com/problems/evaluate-division/
package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	paths := make(map[string]map[string]float64)
	for e, eq := range equations {
		n, d := eq[0], eq[1]
		if _, ok := paths[n]; !ok {
			paths[n] = make(map[string]float64)
			paths[n][n] = 1
		}
		paths[n][d] = values[e]
		if _, ok := paths[d]; !ok {
			paths[d] = make(map[string]float64)
			paths[d][d] = 1
		}
		paths[d][n] = 1 / values[e]
	}

	for from, children := range paths {
		for child := range children {
			for to := range paths[child] {
				if _, ok := paths[from][to]; !ok {
					paths[from][to] = paths[from][child] * paths[child][to]
				}
				if _, ok := paths[to][from]; !ok {
					paths[to][from] = 1 / paths[from][to]
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		n, d := query[0], query[1]
		if val, ok := paths[n][d]; ok {
			res[i] = val
		} else {
			res[i] = -1
		}
	}
	return res
}

func main() {}
