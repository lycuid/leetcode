// https://leetcode.com/problems/3sum-closest/

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	actualSum := sum(nums[:3])
	if l > 3 {
		for i := 0; i < l-2; i++ {
			j := i + 1
			k := l - 1
			for j < k {
				currentSum := nums[i] + nums[j] + nums[k]
				if min(abs(target-actualSum), abs(target-currentSum)) == abs(target-currentSum) {
					actualSum = currentSum
				}

				if currentSum > target {
					k--
				} else if currentSum < target {
					j++
				} else {
					return currentSum
				}
			}
		}
	}
	return actualSum
}

func main() {}
