// https://leetcode.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition/
package main

import "sort"

const MOD = 1e9 + 7

func numSubseq(nums []int, target int) (ret int) {
	sort.Ints(nums)
	pow2 := make([]int, len(nums))
	pow2[0] = 1
	for i := 1; i < len(pow2); i++ {
		pow2[i] = (pow2[i-1] << 1) % MOD
	}
	for l, r := 0, len(nums)-1; l <= r; {
		if nums[l]+nums[r] <= target {
			ret, l = (ret+pow2[r-l])%MOD, l+1
		} else {
			r--
		}
	}
	return ret
}

func main() {}
