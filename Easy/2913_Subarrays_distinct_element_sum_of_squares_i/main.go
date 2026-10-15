// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i/
package main

func sumCounts(nums []int) (res int) {
	var highest int
	for _, num := range nums {
		highest = max(highest, num)
	}
	cache := make([][]int, len(nums)+1)
	for i := range cache {
		cache[i] = make([]int, highest+1)
	}
	for i, num := range nums {
		for j := range cache[i+1] {
			cache[i+1][j] = cache[i][j]
		}
		cache[i+1][num]++
	}
	for i := 1; i < len(cache); i++ {
		for j := 1; j <= i; j++ {
			var count int
			for k := range cache[i] {
				if cache[i][k]-cache[j-1][k] > 0 {
					count++
				}
			}
			res += count * count
		}
	}
	return res
}

func main() {}
