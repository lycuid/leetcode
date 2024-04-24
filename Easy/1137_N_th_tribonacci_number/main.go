// https://leetcode.com/problems/n-th-tribonacci-number/
package main

func tribonacci(n int) int {
	cache := [3]int{0, 1, 1}
	for i := 3; i <= n; i++ {
		cache[i%3] = cache[(i-1)%3] + cache[(i-2)%3] + cache[(i-3)%3]
	}
	return cache[n%3]
}

func main() {}
