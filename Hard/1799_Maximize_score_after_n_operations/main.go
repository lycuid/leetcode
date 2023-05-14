// https://leetcode.com/problems/maximize-score-after-n-operations/
package main

import "sort"

type Tuple struct{ i, j, value int }

func Gcd(i, j int) int {
	for j != 0 {
		i, j = j, i%j
	}
	return i
}

func Aux(n int, cache []Tuple, used []bool) (ret int) {
	if n > 0 {
		for i := 0; i < len(cache); i++ {
			if t := cache[i]; !used[t.i] && !used[t.j] {
				used[t.i], used[t.j] = true, true
				if m := t.value*n + Aux(n-1, cache[i+1:], used); m > ret {
					ret = m
				}
				used[t.i], used[t.j] = false, false
			}
		}
	}
	return ret
}

func maxScore(nums []int) (ret int) {
	var cache []Tuple
	n := len(nums)
	for i := range nums {
		for j := i + 1; j < n; j++ {
			cache = append(cache, Tuple{i, j, Gcd(nums[i], nums[j])})
		}
	}
	sort.Slice(cache, func(i, j int) bool {
		return cache[i].value > cache[j].value
	})
	return Aux(n/2, cache, make([]bool, n))
}

func main() {}
