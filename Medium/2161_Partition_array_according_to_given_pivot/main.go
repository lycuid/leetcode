// https://leetcode.com/problems/partition-array-according-to-given-pivot/
package main

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = pivot
	}
	i, j := 0, len(res)-1
	for ; len(nums) > 0; nums = nums[1:] {
		switch num := nums[0]; {
		case num < pivot:
			res[i], i = num, i+1
		case num > pivot:
			res[j], j = num, j-1
		}
	}
	for l, r := j+1, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}

func main() {}
