// https://leetcode.com/problems/find-the-town-judge/
package main

func findJudge(n int, trust [][]int) int {
	cache := make([][2]int, n+1)
	for i := range trust {
		cache[trust[i][0]][0]++
		cache[trust[i][1]][1]++
	}
	for i := 1; i <= n; i++ {
		if cache[i][0] == 0 && cache[i][1] == n-1 {
			return i
		}
	}
	return -1
}

func main() {}
