// https://leetcode.com/problems/find-center-of-star-graph/
package main

func findCenter(edges [][]int) int {
	cache := make([]bool, len(edges)+2)
	for _, edge := range edges {
		if cache[edge[0]] {
			return edge[0]
		}
		if cache[edge[1]] {
			return edge[1]
		}
		cache[edge[0]], cache[edge[1]] = true, true
	}
	return 0
}

func main() {}
