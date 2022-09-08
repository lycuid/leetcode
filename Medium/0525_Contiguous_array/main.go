// https://leetcode.com/problems/contiguous-array/
package main

func findMaxLength(nums []int) (ret int) {
	index, sum := map[int]int{0: -1}, 0
	for i, num := range nums {
		if sum++; num == 0 {
			sum -= 2
		}
		if j, found := index[sum]; found {
			if i-j > ret {
				ret = i - j
			}
		} else {
			index[sum] = i
		}
	}
	return ret
}

func main() {}
