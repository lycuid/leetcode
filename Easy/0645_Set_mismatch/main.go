// https://leetcode.com/problems/set-mismatch/
package main

func findErrorNums(nums []int) []int {
	var dup, sum int
	found, n := make([]bool, len(nums)), len(nums)
	for _, num := range nums {
		if found[num-1] {
			dup = num
		}
		found[num-1], sum = true, sum+num
	}
	return []int{dup, ((n * (n + 1)) / 2) - sum + dup}
}

func main() {}
