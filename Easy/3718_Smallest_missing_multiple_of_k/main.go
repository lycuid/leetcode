// https://leetcode.com/problems/smallest-missing-multiple-of-k/
package main

func missingMultiple(nums []int, k int) (res int) {
	var high int
	for _, num := range nums {
		high = max(high, num)
	}
	cache := make([]bool, high+k+1)
	for _, num := range nums {
		cache[num] = true
	}
	for res = k; cache[res]; {
		res += k
	}
	return res
}

func main() {}
