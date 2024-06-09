// https://leetcode.com/problems/subarray-sums-divisible-by-k/
package main

func subarraysDivByK(nums []int, k int) (count int) {
	cache := map[int]int{0: 1}
	for i, sum := 0, 0; i < len(nums); i++ {
		sum = (sum + nums[i]) % k
		key := (sum + k) % k
		count += cache[key]
		cache[key]++
	}
	return count
}

func main() {}
