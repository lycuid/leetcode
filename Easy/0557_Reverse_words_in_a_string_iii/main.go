// https://leetcode.com/problems/reverse-words-in-a-string-iii/
package main

func reverseWords(s string) string {
	ret := []byte(s)
	for i, j := 0, 0; i < len(ret); i++ {
		for j = i; i < len(ret) && ret[i] != ' '; i++ {
		}
		for k := i - 1; j < k; j, k = j+1, k-1 {
			ret[j], ret[k] = ret[k], ret[j]
		}
	}
	return string(ret)
}

func main() {}
