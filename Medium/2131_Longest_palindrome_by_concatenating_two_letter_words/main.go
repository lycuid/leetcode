// https: //leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
package main

func longestPalindrome(words []string) (ret int) {
	freq, middle := [26][26]int{}, 0
	for _, w := range words {
		if y, x := w[0]-'a', w[1]-'a'; freq[x][y] > 0 {
			if freq[x][y], ret = freq[x][y]-1, ret+4; y == x {
				middle--
			}
		} else {
			if freq[y][x]++; y == x {
				middle++
			}
		}
	}
	if middle > 0 {
		ret += 2
	}
	return ret
}

func main() {}
