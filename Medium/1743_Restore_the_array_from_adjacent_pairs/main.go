// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/
package main

func restoreArray(adjacentPairs [][]int) (ret []int) {
	adj := make(map[int][]int)
	for _, pair := range adjacentPairs {
		u, v := pair[0], pair[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	prev, next := adjacentPairs[0][0], adjacentPairs[0][1]
	for key, value := range adj {
		if len(value) == 1 {
			next = key
			break
		}
	}
	ret, prev, next = []int{next}, next, adj[next][0]
	for {
		if ret = append(ret, next); len(adj[next]) == 1 {
			break
		}
		if adj[next][0] == prev {
			prev, next = next, adj[next][1]
		} else {
			prev, next = next, adj[next][0]
		}
	}
	return ret
}

func main() {}
