// https://leetcode.com/problems/largest-number/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func Less(s1, s2 string) bool {
	if n, m := len(s1), len(s2); n != m {
		return n <= m
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		return s1[i] <= s2[i]
	}
	return true
}

func largestNumber(nums []int) (ret string) {
	cache := make([]string, len(nums))
	for i, num := range nums {
		cache[i] = fmt.Sprintf("%d", num)
	}
	sort.Slice(cache, func(i, j int) bool {
		return Less(cache[j]+cache[i], cache[i]+cache[j])
	})
	if n := strings.TrimLeft(strings.Join(cache, ""), "0"); len(n) > 0 {
		return n
	}
	return "0"
}

func main() {}
