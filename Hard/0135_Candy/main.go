// https://leetcode.com/problems/candy/
package main

func candy(ratings []int) int {
	cache, n := make([]int, len(ratings)), len(ratings)
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			cache[i] = cache[i-1] + 1
		}
	}
	ret := cache[n-1] + 1
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && cache[i] < cache[i+1]+1 {
			cache[i] = cache[i+1] + 1
		}
		ret += cache[i] + 1
	}
	return ret
}

func main() {}
