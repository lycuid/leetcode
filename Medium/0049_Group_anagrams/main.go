// https://leetcode.com/problems/group-anagrams/
package main

type Freq struct{ inner [26]int }

func groupAnagrams(strs []string) (ret [][]string) {
	cache := make(map[Freq][]string)
	for _, str := range strs {
		var freq Freq
		for _, ch := range str {
			freq.inner[ch-'a']++
		}
		cache[freq] = append(cache[freq], str)
	}
	for _, value := range cache {
		ret = append(ret, value)
	}
	return ret
}

func main() {}
