// https://leetcode.com/problems/palindromic-substrings/
package main

func Count(s string, i, j int) (count int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		count++
	}
	return count
}

func countSubstrings(s string) (count int) {
	for i := range s {
		count += Count(s, i, i) + Count(s, i-1, i)
	}
	return count
}

func main() {}
