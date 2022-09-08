// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
package main

func minDeletions(s string) (ret int) {
	freq, taken := make([]int, 26), make([]bool, len(s)+1)
	for _, ch := range s {
		freq[ch-'a']++
	}
	for _, f := range freq {
		for ; f > 0 && taken[f]; f-- {
			ret++
		}
		taken[f] = true
	}
	return ret
}

func main() {}
