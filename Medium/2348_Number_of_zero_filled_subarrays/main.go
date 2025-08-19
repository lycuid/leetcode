// https://leetcode.com/problems/number-of-zero-filled-subarrays/
package main

func zeroFilledSubarray(nums []int) (ret int64) {
	var count int64
	for _, num := range nums {
		if num == 0 {
			count++
		} else {
			count = 0
		}
		ret += count
	}
	return ret
}

func main() {}
