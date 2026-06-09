// https://leetcode.com/problems/maximum-total-subarray-value-i/
package main

func maxTotalValue(nums []int, k int) int64 {
	_min, _max := nums[0], nums[0]
	for _, num := range nums {
		_min, _max = min(_min, num), max(_max, num)
	}
	return int64(k * (_max - _min))
}

func main() {}
