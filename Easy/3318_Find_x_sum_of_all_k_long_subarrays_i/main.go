// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/
package main

import "sort"

func findXSum(nums []int, k int, x int) []int {
	var (
		n     = len(nums)
		xsum  = make([]int, 0, n-k+1)
		freq  = make(map[int]int)
		cache = make([]int, 0, k)
	)

	for i := 0; i < k-1; i++ {
		freq[nums[i]]++
	}
	for i := 0; i+k-1 < n; i++ {
		freq[nums[i+k-1]]++
		for num := range freq {
			cache = append(cache, num)
		}

		sort.Slice(cache, func(i, j int) bool {
			if freq[cache[i]] == freq[cache[j]] {
				return cache[i] > cache[j]
			}
			return freq[cache[i]] > freq[cache[j]]
		})

		var sum int
		for _, num := range cache[:min(x, len(cache))] {
			sum += num * freq[num]
		}
		xsum, cache = append(xsum, sum), cache[:0]

		if freq[nums[i]]--; freq[nums[i]] == 0 {
			delete(freq, nums[i])
		}
	}
	return xsum
}

func main() {}
