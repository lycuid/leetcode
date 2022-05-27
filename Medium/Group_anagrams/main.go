// https://leetcode.com/problems/group-anagrams/
package main

type AlphaCount struct {
	inner [26]int
}

func groupAnagrams(strs []string) (sets [][]string) {
	grams := make(map[AlphaCount][]string)
	for _, str := range strs {
		count := AlphaCount{}
		for _, ch := range str {
			count.inner[ch-'a']++
		}
		grams[count] = append(grams[count], str)
	}
	for _, set := range grams {
		sets = append(sets, set)
	}
	return sets
}

func main() {}
