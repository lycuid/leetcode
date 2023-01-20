// https://leetcode.com/problems/non-decreasing-subsequences/
package main

func ListEq(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func In(item []int, items [][]int) bool {
	for i := range items {
		if len(item) == len(items[i]) && ListEq(item, items[i]) {
			return true
		}
	}
	return false
}

func Aux(i int, nums []int, cache [][][]int) {
	if len(cache[i]) > 0 {
		return
	}
	for j := i + 1; j < len(nums); j++ {
		if nums[i] <= nums[j] {
			Aux(j, nums, cache)
			cache[i] = append(cache[i], []int{nums[i], nums[j]})
			for _, seq := range cache[j] {
				cache[i] = append(cache[i], append([]int{nums[i]}, seq...))
			}
		}
	}
}

func findSubsequences(nums []int) (ret [][]int) {
	cache := make([][][]int, len(nums))
	for i := range nums {
		Aux(i, nums, cache)
	}
	for i := range cache {
		for j := range cache[i] {
			if !In(cache[i][j], ret) {
				ret = append(ret, cache[i][j])
			}
		}
	}
	return ret
}

func main() {}
