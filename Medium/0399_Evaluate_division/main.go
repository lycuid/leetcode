// https://leetcode.com/problems/evaluate-division/
package main

type Solution struct {
	edges    map[string]map[string]float64
	on_stack map[string]bool
}

func MakeSolution(equations [][]string, values []float64) *Solution {
	edges := make(map[string]map[string]float64)
	for i, eqn := range equations {
		from, to := eqn[0], eqn[1]
		if _, found := edges[from]; !found {
			edges[from] = make(map[string]float64)
		}
		if _, found := edges[to]; !found {
			edges[to] = make(map[string]float64)
		}
		edges[from][to], edges[to][from] = values[i], 1/values[i]
	}
	return &Solution{edges, make(map[string]bool)}
}

func (this *Solution) Query(from, to string) float64 {
	if val, found := this.edges[from][to]; found {
		return val
	}
	if val, found := this.edges[to][from]; found {
		return 1 / val
	}
	this.on_stack[from] = true
	for child, value := range this.edges[from] {
		if !this.on_stack[child] {
			if dist := this.Query(child, to); dist != -1 {
				this.edges[from][to] = value * dist
				break
			}
		}
	}
	this.on_stack[from] = false
	if value, found := this.edges[from][to]; found {
		return value
	}
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (ret []float64) {
	solution := MakeSolution(equations, values)
	for _, query := range queries {
		ret = append(ret, solution.Query(query[0], query[1]))
	}
	return ret
}

func main() {}
