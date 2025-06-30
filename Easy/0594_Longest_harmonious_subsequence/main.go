// https://leetcode.com/problems/longest-harmonious-subsequence/
package main

func findLHS(nums []int) (size int) {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	for num, freq := range cache {
		if value, found := cache[num+1]; found {
			size = max(size, value+freq)
		}
		if value, found := cache[num-1]; found {
			size = max(size, value+freq)
		}
	}
	return size
}

func main() {}
