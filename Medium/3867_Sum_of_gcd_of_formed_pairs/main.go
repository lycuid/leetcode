// https://leetcode.com/problems/sum-of-gcd-of-formed-pairs/
package main

import "sort"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcdSum(nums []int) (res int64) {
	var (
		n = len(nums)
		m int
	)
	prefix := make([]int, n)
	for i, num := range nums {
		m = max(m, num)
		prefix[i] = gcd(num, m)
	}
	sort.Ints(prefix)
	for i := 0; i < n/2; i++ {
		res += int64(gcd(prefix[i], prefix[n-1-i]))
	}
	return res
}

func main() {}
