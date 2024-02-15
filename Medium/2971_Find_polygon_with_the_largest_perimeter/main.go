// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/
package main

import "sort"

func largestPerimeter(nums []int) (ret int64) {
	sort.Ints(nums)
	nums[1], ret = nums[0]+nums[1], -1
	for i := 2; i < len(nums); i++ {
		if nums[i] < nums[i-1] && ret < int64(nums[i]+nums[i-1]) {
			ret = int64(nums[i] + nums[i-1])
		}
		nums[i] += nums[i-1]
	}
	return ret
}

func main() {}
