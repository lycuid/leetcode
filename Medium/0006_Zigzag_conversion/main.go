// https://leetcode.com/problems/zigzag-conversion/
package main

func getDistances(n int) []int {
	distances := make([]int, n)
	distances[n-1] = 2*n - 2
	for i := 0; i < n-1; i++ {
		distances[i] = 2*n - (2 * (i + 1))
	}
	return distances
}

func convert(s string, n int) string {
	if n <= 1 {
		return s
	}
	ret, distances := make([]byte, len(s)), getDistances(n)
	for i, n, x := 0, len(distances), 0; i < n; i++ {
		for j, pointer := 0, i; pointer < len(s); j++ {
			ret[x], x = s[pointer], x+1
			if j%2 == 0 {
				pointer += distances[i]
			} else {
				pointer += distances[n-i-1]
			}
		}
	}
	return string(ret)
}

func main() {}
