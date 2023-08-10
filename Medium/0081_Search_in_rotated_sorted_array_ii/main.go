// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
package main

func search(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func main() {}
