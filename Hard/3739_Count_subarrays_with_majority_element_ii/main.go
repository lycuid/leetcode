// https://leetcode.com/problems/count-subarrays-with-majority-element-ii/
package main

func countMajoritySubarrays(nums []int, target int) (count int64) {
	var (
		n      = len(nums)
		cache  = make([]int64, 2*n+1)
		presum int64
	)
	cache[n] = 1
	for _, num := range nums {
		switch num {
		case target:
			presum = presum + cache[n]
			n++
		default:
			n--
			presum = presum - cache[n]
		}
		cache[n]++
		count += presum
	}
	return count
}

func main() {}
