// https://leetcode.com/problems/contains-duplicate-ii/
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	index := make(map[int]int)
	for i, num := range nums {
		if j, found := index[num]; found && i-j <= k {
			return true
		}
		index[num] = i
	}
	return false
}

func main() {}
