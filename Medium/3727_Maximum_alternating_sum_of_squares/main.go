// https://leetcode.com/problems/maximum-alternating-sum-of-squares/
package main

func maxAlternatingSum(nums []int) (res int64) {
	var highest int
	for i := range nums {
		nums[i] = max(-nums[i], nums[i])
		highest = max(highest, nums[i])
	}
	cache := make([]int, highest+1)
	for _, num := range nums {
		cache[num]++
	}
	l := len(nums)
	for have := l>>1 + l&1; highest > 0; highest-- {
		if cache[highest] > 0 {
			take := min(have, cache[highest])
			res += int64(highest*highest) * int64(take)
			if have -= take; have == 0 {
				cache[highest] -= take
				break
			}
		}
	}
	for have := l >> 1; have > 0 && highest > 0; highest-- {
		if cache[highest] > 0 {
			take := min(have, cache[highest])
			res -= int64(highest*highest) * int64(take)
			have -= take
		}
	}
	return res
}

func main() {}
