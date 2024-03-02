// https://leetcode.com/problems/squares-of-a-sorted-array/
package main

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	for i, j, k := 0, len(nums)-1, len(ret)-1; i <= j; {
		if si, sj := nums[i]*nums[i], nums[j]*nums[j]; si > sj {
			ret[k], k, i = si, k-1, i+1
		} else {
			ret[k], k, j = sj, k-1, j-1
		}
	}
	return ret
}

func main() {}
