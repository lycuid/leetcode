// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
package main

func minDeletions(s string) (ret int) {
	freq, occupied := make([]int, 26), make([]bool, len(s)+1)
	for _, ch := range s {
		freq[ch-'a']++
	}
	for _, f := range freq {
		for ; f > 0 && occupied[f]; f-- {
			ret++
		}
		occupied[f] = true
	}
	return ret
}

func main() {}
