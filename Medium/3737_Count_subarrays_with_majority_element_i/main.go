// https://leetcode.com/problems/count-subarrays-with-majority-element-i/
package main

func countMajoritySubarrays(nums []int, target int) (res int) {
	cache := make([]int, len(nums)+1)
	for i, num := range nums {
		if cache[i+1] += cache[i]; num == target {
			cache[i+1]++
		}
	}
	for i := 1; i < len(cache); i++ {
		for j := i; j < len(cache); j++ {
			if cache[j]-cache[i-1] > (j-i+1)>>1 {
				res++
			}
		}
	}
	return res
}

func main() {}
