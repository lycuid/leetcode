// https://leetcode.com/problems/orderly-queue/
package main

import "sort"

func orderlyQueue(s string, k int) string {
	if tmp := s; k == 1 {
		for i := range s {
			if str := s[i:] + s[:i]; str < tmp {
				tmp = str
			}
		}
		return tmp
	}
	ret := []byte(s)
	sort.Slice(ret, func(i, j int) bool { return ret[i] < ret[j] })
	return string(ret)
}

func main() {}
