// https://leetcode.com/problems/3sum/

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

func sum(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return sum
}

func threeSum(nums []int) (threeSums [][]int) {
	sort.Ints(nums)
	l := len(nums)

	if l <= 3 {
		if sum(nums) == 0 && l == 3 {
			return [][]int{nums}
		}
		return [][]int{}
	}

	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
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

func main() {}
