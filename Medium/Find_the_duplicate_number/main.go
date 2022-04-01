// https://leetcode.com/problems/find-the-duplicate-number/
package main

func findDuplicate(nums []int) int {
	seen := make([]bool, len(nums))
	for _, n := range nums {
		if seen[n] {
			return n
		}
		seen[n] = true
	}
	return 0
}

func main() {}
