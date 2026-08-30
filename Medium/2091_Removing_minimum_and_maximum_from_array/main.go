// https://leetcode.com/problems/removing-minimum-and-maximum-from-array/
package main

func minimumDeletions(nums []int) int {
	minn, maxn := nums[0], nums[0]
	mini, maxi := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < minn {
			minn, mini = nums[i], i
		}
		if nums[i] > maxn {
			maxn, maxi = nums[i], i
		}
	}
	l, r := min(mini, maxi), max(mini, maxi)
	return min(r+1, len(nums)-l, len(nums)-r+l+1)
}

func main() {}
