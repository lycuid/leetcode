// https://leetcode.com/problems/find-original-array-from-doubled-array/
package main

import "sort"

func findOriginalArray(changed []int) (ret []int) {
	if n := len(changed); n%2 == 0 {
		var cache [1e5 + 1]int
		sort.Ints(changed)
		for _, val := range changed {
			if val%2 == 1 || cache[val/2] == 0 {
				cache[val], ret = cache[val]+1, append(ret, val)
			} else {
				cache[val/2]--
			}
		}
		if len(ret) != n/2 {
			ret = nil
		}
	}
	return ret
}

func main() {}
