// https://leetcode.com/problems/minimum-sum-of-squared-difference/
package main

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) (res int64) {
	var highest int
	for i := range nums1 {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		highest = max(highest, diff)
	}
	cache := make([]int64, highest+1)
	for i := range nums1 {
		diff := nums1[i] - nums2[i]
		if diff < 0 {
			diff = -diff
		}
		cache[diff]++
	}
	total := int64(k1 + k2)
	for i := highest; i > 0; i-- {
		if count := cache[i]; count > 0 {
			carry := min(count, total)
			total = max(0, total-count)
			cache[i-1] = cache[i-1] + carry
			res += int64(i*i) * (count - carry)
		}
	}
	return res
}

func main() {}
