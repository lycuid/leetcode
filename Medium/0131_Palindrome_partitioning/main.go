// https://leetcode.com/problems/palindrome-partitioning/
package main

import "fmt"

func IsPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func Aux(s string, start int, cache [][][]string) [][]string {
	for i := start; i < len(s); i++ {
		if IsPalindrome(s[start : i+1]) {
			if cache[i+1] == nil {
				Aux(s, i+1, cache)
			}
			for _, part := range cache[i+1] {
				cache[start] = append(cache[start], append([]string{s[start : i+1]}, part...))
			}
		}
	}
	return cache[start]
}

func partition(s string) (ret [][]string) {
	cache := make([][][]string, len(s)+1)
	cache[len(s)] = [][]string{{}}
	return Aux(s, 0, cache)
}

func main() {
	fmt.Println(partition("a"))
	fmt.Println(partition("aab"))
	fmt.Println(partition("aaa"))
}
