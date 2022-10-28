// https://leetcode.com/problems/group-anagrams/
package main

import "sort"

func SortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}

func groupAnagrams(strs []string) (ret [][]string) {
	cache := make(map[string][]string)
	for _, str := range strs {
		sorted := SortString(str)
		cache[sorted] = append(cache[sorted], str)
	}
	for _, val := range cache {
		ret = append(ret, val)
	}
	return ret
}

func main() {}
