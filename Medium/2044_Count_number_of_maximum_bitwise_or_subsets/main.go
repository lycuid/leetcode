// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
package main

func permute(nums []int) (out map[int]int) {
	if out = make(map[int]int); len(nums) > 0 {
		head, tail := nums[0], nums[1:]
		out[head] = 1
		for num, count := range permute(tail) {
			out[num] += count
			out[head|num] += count
		}
	}
	return out
}

func countMaxOrSubsets(nums []int) (maxcount int) {
	var maxnum int
	for num, count := range permute(nums) {
		if num >= maxnum {
			maxcount = max(maxcount, count)
		}
	}
	return maxcount
}

func main() {}
