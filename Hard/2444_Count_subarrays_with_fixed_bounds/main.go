// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
package main

func countSubarrays(nums []int, minK int, maxK int) (count int64) {
	mini, maxi := -1, -1
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] < minK || nums[j] > maxK {
			mini, maxi, i = -1, -1, j+1
		}
		if nums[j] == minK {
			mini = j
		}
		if nums[j] == maxK {
			maxi = j
		}
		index := mini
		if index > maxi {
			index = maxi
		}
		if index >= i {
			count += int64(index - i + 1)
		}
	}
	return count
}

func main() {}
