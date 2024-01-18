// https://leetcode.com/problems/climbing-stairs/
package main

func climbStairs(n int) int {
	cache := make([]int, n+1)
	cache[0], cache[1] = 1, 1
	for i := 2; i < len(cache); i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}

func main() {}
