// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/
package main

func countSubarrays(nums []int, k int) (count int64) {
	max, n := nums[0], len(nums)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	for i, j, f := 0, 0, 0; j < n; j++ {
		if nums[j] == max {
			f++
		}
		for ; i <= j && f >= k; i++ {
			if nums[i] == max {
				f--
			}
		}
		count += int64(j - i + 1)
	}
	return int64(n*(n+1))/2 - count
}

func main() {}
