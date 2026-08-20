// https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
package main

func resultArray(nums []int) []int {
	res := make([]int, len(nums))
	res[0], res[len(nums)-1] = nums[0], nums[1]
	l, r := 1, len(nums)-2
	for _, num := range nums[2:] {
		if res[l-1] > res[r+1] {
			res[l], l = num, l+1
		} else {
			res[r], r = num, r-1
		}
	}
	for r = len(res) - 1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}

func main() {}
