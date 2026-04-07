// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/
package main

func minimumOperations(nums []int) (count int) {
	for _, num := range nums {
		count += min(1, num%3)
	}
	return count
}

func main() {}
