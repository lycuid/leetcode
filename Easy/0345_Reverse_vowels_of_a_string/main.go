// https://leetcode.com/problems/reverse-vowels-of-a-string/
package main

func reverseVowels(s string) string {
	b, IsVowel := []byte(s), func(b byte) bool {
		return b == 'a' || b == 'A' || b == 'e' || b == 'E' || b == 'i' ||
				b == 'I' || b == 'o' || b == 'O' || b == 'u' || b == 'U'
	}
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		for ; i < j && !IsVowel(b[i]); i++ {
		}
		for ; i < j && !IsVowel(b[j]); j-- {
		}
		if IsVowel(b[i]) && IsVowel(b[j]) {
			b[i], b[j] = b[j], b[i]
		}
	}
	return string(b)
}

func main() {}
