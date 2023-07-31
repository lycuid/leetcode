// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
package main

func minimumDeleteSum(s1, s2 string) int {
	n, m, sum := len(s1), len(s2), 0
	prev, curr := make([]int, n+1), make([]int, n+1)
	for _, ch := range s1 {
		sum += int(ch)
	}
	for _, ch := range s2 {
		sum += int(ch)
	}
	for i := 1; i <= m; i++ {
		prev, curr = curr, prev
		for j := 1; j <= n; j++ {
			if curr[j] = curr[j-1]; curr[j] < prev[j] {
				curr[j] = prev[j]
			}
			if s1[j-1] == s2[i-1] {
				if n := int(s1[j-1]) + prev[j-1]; curr[j] < n {
					curr[j] = n
				}
			}
		}
	}
	return sum - curr[n]*2
}

func main() {}
