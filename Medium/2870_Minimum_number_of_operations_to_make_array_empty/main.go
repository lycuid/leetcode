// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/
package main

func minOperations(nums []int) (count int) {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}
	for _, freq := range cache {
		if freq%3 == 0 {
			count += freq / 3
		} else if n := freq - 2; n >= 0 && n%3 == 0 {
			count += n/3 + 1
		} else if n := freq - 4; n >= 0 && n%3 == 0 {
			count += n/3 + 2
		} else if freq%2 == 0 {
			count += freq / 2
		} else {
			return -1
		}
	}
	return count
}

func main() {}
