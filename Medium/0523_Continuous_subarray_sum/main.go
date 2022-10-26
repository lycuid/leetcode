// https://leetcode.com/problems/continuous-subarray-sum/
package main

func checkSubarraySum(nums []int, k int) bool {
	for i, sum, prefix := 0, 0, make(map[int]int); i < len(nums); i++ {
		if sum = (sum + nums[i]) % k; sum == 0 && i > 0 {
			return true
		}
		if index, found := prefix[sum]; found {
			if i-index > 1 {
				return true
			}
		} else {
			prefix[sum] = i
		}
	}
	return false
}

func main() {}
