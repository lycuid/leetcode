// https://leetcode.com/problems/climbing-stairs-ii/
package main

func climbStairs(n int, costs []int) int {
	cache, res := make([]int, 4), 0
	for _, cost := range costs {
		res = (res + 1) & 0b11
		cache[res] = min(
			cache[(res+1)&0b11]+cost+9,
			cache[(res+2)&0b11]+cost+4,
			cache[(res+3)&0b11]+cost+1,
		)
	}
	return cache[res]
}

func main() {}
