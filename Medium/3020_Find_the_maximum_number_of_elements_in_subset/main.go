// https://leetcode.com/problems/find-the-maximum-number-of-elements-in-subset/
package main

func maximumLength(nums []int) (res int) {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}

	var solve func(int) int
	solve = func(num int) (res int) {
		switch cache[num] {
		case 0:
			res = -1
		case 1:
			res = 1
		default:
			cache[num] -= 2
			res = 2 + solve(num*num)
			cache[num] += 2
		}
		return res
	}

	for num := range cache {
		res = max(res, solve(num))
	}
	return res
}

func main() {}
