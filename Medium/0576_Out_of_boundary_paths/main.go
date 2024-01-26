// https://leetcode.com/problems/out-of-boundary-paths/
package main

func findPaths(m, n, moves, start_r, start_c int) (ret int) {
	if moves == 0 {
		return 0
	}
	const MOD = 1e9 + 7
	OpenEdges := func(r, c int) (ans int) {
		if r == 0 {
			ans++
		}
		if r == m-1 {
			ans++
		}
		if c == 0 {
			ans++
		}
		if c == n-1 {
			ans++
		}
		return ans
	}
	var path [2][][]int
	for i := range path {
		path[i] = make([][]int, m)
		for j := range path[i] {
			path[i][j] = make([]int, n)
		}
	}
	path[0][start_r][start_c] = 1
	ret = OpenEdges(start_r, start_c)
	for prev, cur := 0, 1; moves > 1; prev, cur, moves = cur, prev, moves-1 {
		for r := range path[cur] {
			for c := range path[cur][r] {
				path[cur][r][c] = 0
				if r > 0 {
					path[cur][r][c] += path[prev][r-1][c]
				}
				if r < m-1 {
					path[cur][r][c] += path[prev][r+1][c]
				}
				if c > 0 {
					path[cur][r][c] += path[prev][r][c-1]
				}
				if c < n-1 {
					path[cur][r][c] += path[prev][r][c+1]
				}
				path[cur][r][c] %= MOD
				ret = (ret + path[cur][r][c]*OpenEdges(r, c)) % MOD
			}
		}
	}
	return ret % MOD
}

func main() {}
