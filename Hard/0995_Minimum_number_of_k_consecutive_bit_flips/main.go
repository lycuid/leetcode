// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/
package main

func minKBitFlips(nums []int, k int) int {
	cache := make([]int, 0, len(nums))
	for i, num := range nums {
		l, r := 0, len(cache)
		for l < r {
			if mid := l + (r-l)/2; i > cache[mid] {
				l = mid + 1
			} else if i < cache[mid] {
				r = mid
			} else {
				l, r = mid, mid
			}
		}
		if (num+len(cache)-l)%2 == 0 {
			if i+k-1 >= len(nums) {
				return -1
			}
			cache = append(cache, i+k-1)
		}
	}
	return len(cache)
}

func main() {}
