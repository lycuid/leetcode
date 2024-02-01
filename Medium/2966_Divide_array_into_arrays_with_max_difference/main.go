// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
package main

import "sort"

func divideArray(nums []int, k int) (ret [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		ret = append(ret, nums[i:i+3])
	}
	return ret
}

func main() {}
