// https://leetcode.com/problems/word-pattern/
package main

func wordPattern(pattern string, s string) bool {
	charOf, strAt, i, j := make(map[string]byte), make(map[byte]string), 0, 0
	for s += " "; i < len(pattern) && j < len(s); i, j = i+1, j+1 {
		k := j
		for ; j < len(s) && s[j] != ' '; j++ {
		}
		ch, str := pattern[i], s[k:j]
		if _, found := charOf[str]; !found {
			charOf[str] = ch
		}
		if _, found := strAt[ch]; !found {
			strAt[ch] = str
		}
		if charOf[str] != ch || strAt[ch] != str {
			return false
		}
	}
	return i == len(pattern) && j == len(s)
}

func main() {}
