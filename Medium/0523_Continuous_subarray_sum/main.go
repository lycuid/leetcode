// https://leetcode.com/problems/continuous-subarray-sum/
package main

func checkSubarraySum(nums []int, k int) bool {
	prefix, sum := make(map[int]int), 0
	for i, num := range nums {
		if sum = (sum + num) % k; sum == 0 && i > 0 {
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
