// https://leetcode.com/problems/custom-sort-string/
package main

import "sort"

func customSortString(order, s string) string {
	var lang [26]int
	for i := range lang {
		lang[i] = 26
	}
	for i, ch := 0, 0; ch < len(order); ch++ {
		lang[order[ch]-'a'], i = i, i+1
	}
	ret := []byte(s)
	sort.Slice(ret, func(i, j int) bool {
		return lang[ret[i]-'a'] < lang[ret[j]-'a']
	})
	return string(ret)
}

func main() {}
