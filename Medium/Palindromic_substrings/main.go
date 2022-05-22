// https://leetcode.com/problems/palindromic-substrings/
package main

func countSubstrings(str string) (count int) {
	for c := range str {
		for _, dj := range []int{0, 1} {
			i, j := c, c+dj
			for i >= 0 && j < len(str) && str[i] == str[j] {
				count++
				i--
				j++
			}
		}
	}
	return count
}

func main() {}
