// https://leetcode.com/problems/4sum/

package main

import "sort"

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, val := range x {
		if val != y[i] {
			return false
		}
	}
	return true
}

func in(arr []int, nested [][]int) (present bool) {
	for i := 0; i < len(nested); i++ {
		if equal(arr, nested[i]) {
			return true
		}
	}
	return false
}

func sum(x []int) (sumVal int) {
	for _, v := range x {
		sumVal += v
	}
	return sumVal
}

func threeSum(nums []int, target int) (threeSums [][]int) {
	l := len(nums)

	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] > target {
				k--
			} else if nums[i]+nums[j]+nums[k] < target {
				j++
			} else {
				temp := []int{nums[i], nums[j], nums[k]}
				if len(threeSums) == 0 || !in(temp, threeSums) {
					threeSums = append(threeSums, temp)
				}
				k--
				j++
			}
		}
	}

	return threeSums
}

func fourSum(nums []int, target int) (fourSums [][]int) {
	sort.Ints(nums)
	l := len(nums)
	if l <= 4 {
		if sum(nums) == target && l == 4 {
			return [][]int{nums}
		}
		return [][]int{}
	}

	for i, num := range nums {
		for _, threeSums := range threeSum(nums[i+1:], target-num) {
			if len(threeSums) > 0 {
				threeSums = append([]int{num}, threeSums...)
				if len(fourSums) == 0 || !in(threeSums, fourSums) {
					fourSums = append(fourSums, threeSums)
				}
			}
		}
	}
	return fourSums
}

func main() {}
