// https://leetcode.com/problems/permutations-ii/
package main

func Equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func In(needle []int, haystack [][]int) bool {
	for _, val := range haystack {
		if Equal(needle, val) {
			return true
		}
	}
	return false
}

func Combinations(index, end int, set, used, nums []int, ret *[][]int) {
	if index < end {
		for i, num := range nums {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			set = append(set, num)
			Combinations(index+1, end, set, used, nums, ret)
			if len(set) == end && !In(set, *ret) {
				tmp := make([]int, end)
				for i := range set {
					tmp[i] = set[i]
				}
				*ret = append(*ret, tmp)
			}
			set = set[:len(set)-1]
			used[i] = 0
		}
	}
}

func permuteUnique(nums []int) (ret [][]int) {
	var set []int
	used := make([]int, len(nums))
	Combinations(0, len(nums), set, used, nums, &ret)
	return ret
}

func main() {}
