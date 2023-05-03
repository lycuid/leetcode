// https://leetcode.com/problems/find-the-difference-of-two-arrays/
package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	ret := [][]int{nil, nil}
	set := [2]map[int]bool{make(map[int]bool), make(map[int]bool)}
	for _, num := range nums1 {
		set[1][num] = true
	}
	for _, num := range nums2 {
		set[0][num] = true
	}
	for num := range set[1] {
		if !set[0][num] {
			ret[0] = append(ret[0], num)
		}
	}
	for num := range set[0] {
		if !set[1][num] {
			ret[1] = append(ret[1], num)
		}
	}
	return ret
}

func main() {}
