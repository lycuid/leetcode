// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/
package main

func countPrefixSuffixPairs(words []string) (count int) {
	aux := func(s1, s2 string) bool {
		return len(s1) <= len(s2) && s1 == s2[:len(s1)] && s1 == s2[len(s2)-len(s1):]
	}
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if aux(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func main() {}
