// https://leetcode.com/problems/perfect-squares/
package main

func numSquares(n int) int {
	cache := []int{0}
	for i := len(cache); i <= n; i++ {
		cache = append(cache, i)
		for j := i - 1; j >= i/2; j-- {
			if k := i - j; k*k == i {
				cache[i] = 1
			} else if m := cache[j] + cache[k]; m < cache[i] {
				cache[i] = m
			}
		}
	}
	return cache[n]
}

func main() {}
