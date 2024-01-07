// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
package main

func numberOfArithmeticSlices(nums []int) (count int) {
	cache := make([]map[int]int, len(nums))
	for i := range cache {
		cache[i] = make(map[int]int)
	}
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			diff := nums[i] - nums[j]
			cache[i][diff] += 1 + cache[j][diff]
			count += cache[j][diff]
		}
	}
	return count
}

func main() {}
