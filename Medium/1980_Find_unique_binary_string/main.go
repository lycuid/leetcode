// https://leetcode.com/problems/find-unique-binary-string/
package main

import "sort"

func findDifferentBinaryString(nums []string) string {
	cache := make([]int, len(nums))
	for i, num := range nums {
		for _, ch := range num {
			cache[i] = (cache[i] << 1) | int(ch-'0')
		}
	}
	sort.Ints(cache)
	var num int
	for ; len(cache) > 0 && num >= cache[0]; cache = cache[1:] {
		if num == cache[0] {
			num++
		}
	}
	ret := make([]byte, 0, len(nums))
	for ; num > 0; num >>= 1 {
		ret = append(ret, byte(num&1)+'0')
	}
	for len(ret) < len(nums) {
		ret = append(ret, '0')
	}
	for i, n := 0, len(ret)-1; i <= n/2; i++ {
		ret[i], ret[n-i] = ret[n-i], ret[i]
	}
	return string(ret)
}

func main() {}
