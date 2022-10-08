// https://leetcode.com/problems/3sum-closest/
package main

import "sort"
import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := math.MaxInt
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if abs(target-sum) < abs(target-ret) {
				ret = sum
			}
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				return sum
			}
		}
	}
	return ret
}

func main() {}
