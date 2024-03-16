// https://leetcode.com/problems/contiguous-array/
package main

func findMaxLength(nums []int) (ret int) {
	cache, sum := map[int]int{0: -1}, 0
	for i, num := range nums {
		if sum++; num == 0 {
			sum -= 2
		}
		if j, found := cache[sum]; found {
			if i-j > ret {
				ret = i - j
			}
		} else {
			cache[sum] = i
		}
	}
	return ret
}

func main() {}
