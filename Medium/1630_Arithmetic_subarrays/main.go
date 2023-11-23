// https://leetcode.com/problems/arithmetic-subarrays/
package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ret := make([]bool, len(l))
mainloop:
	for i := range ret {
		slice := make([]int, r[i]-l[i]+1)
		for j, k := l[i], 0; j <= r[i]; j, k = j+1, k+1 {
			slice[k] = nums[j]
		}
		sort.Ints(slice)
		for p, interval := 1, slice[1]-slice[0]; p < len(slice); p++ {
			if slice[p]-slice[p-1] != interval {
				ret[i] = false
				continue mainloop
			}
		}
		ret[i] = true
	}
	return ret
}

func main() {}
