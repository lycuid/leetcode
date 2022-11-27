// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
package main

func numberOfArithmeticSlices(nums []int) (ret int) {
	cache := make([]map[int]int, len(nums))
	for i := range cache {
		cache[i] = make(map[int]int)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			cache[i][diff] += (cache[j][diff] + 1)
			ret += cache[j][diff]
		}
	}
	return ret
}

func main() {}
