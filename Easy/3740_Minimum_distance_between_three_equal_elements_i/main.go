// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
package main

func minimumDistance(nums []int) int {
	cache := make([][2]int, len(nums)+1)
	for i := range cache {
		cache[i] = [2]int{-1, -1}
	}
	res := len(nums)*2 + 1
	for i, num := range nums {
		if c := cache[num]; c[0] != -1 {
			res = min(res, c[1]-c[0]+i-c[1]+i-c[0])
		}
		cache[num][0], cache[num][1] = cache[num][1], i
	}
	if res > len(nums)*2 {
		return -1
	}
	return res
}

func main() {}
