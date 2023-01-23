// https://leetcode.com/problems/find-the-town-judge/
package main

func findJudge(n int, trust [][]int) int {
	in, out := make([]int, n+1), make([]int, n+1)
	for _, t := range trust {
		in[t[1]], out[t[0]] = in[t[1]]+1, out[t[0]]+1
	}
	for i := 1; i <= n; i++ {
		if in[i] == n-1 && out[i] == 0 {
			return i
		}
	}
	return -1
}

func main() {}
