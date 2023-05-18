// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/
package main

func findSmallestSetOfVertices(n int, edges [][]int) (ret []int) {
	visited := make([]bool, n)
	for _, e := range edges {
		visited[e[1]] = true
	}
	for i := range visited {
		if !visited[i] {
			ret = append(ret, i)
		}
	}
	return ret
}

func main() {}
