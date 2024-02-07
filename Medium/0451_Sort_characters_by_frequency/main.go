// https://leetcode.com/problems/sort-characters-by-frequency/
package main

import "sort"

func frequencySort(s string) string {
	var freq [128]int
	for _, ch := range s {
		freq[ch]++
	}
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		if freq[str[i]] == freq[str[j]] {
			return str[i] < str[j]
		}
		return freq[str[i]] > freq[str[j]]
	})
	return string(str)
}

func main() {}
