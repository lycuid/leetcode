// https://leetcode.com/problems/max-number-of-k-sum-pairs/
package main

func maxOperations(nums []int, k int) (sum int) {
	counts := make(map[int]int)
	for _, num := range nums {
		comp := k - num
		if counts[comp] > 0 {
			counts[comp]--
			sum++
		} else {
			counts[num]++
		}
	}
	return sum
}

func main() {}
