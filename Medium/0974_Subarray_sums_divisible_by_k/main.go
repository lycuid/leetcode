// https://leetcode.com/problems/subarray-sums-divisible-by-k/
package main

func subarraysDivByK(nums []int, k int) (ret int) {
	pre, mod := make([]int, len(nums)+1), make([]int, k)
	for i := 1; i < len(pre); i++ {
		pre[i] = pre[i-1] + nums[i-1]
	}
	for _, sum := range pre {
		mod[(sum%k+k)%k]++
	}
	for _, num := range mod {
		ret += (num*num - num) / 2
	}
	return ret
}

func main() {}
