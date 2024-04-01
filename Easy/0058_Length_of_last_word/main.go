// https://leetcode.com/problems/length-of-last-word/
package main

func lengthOfLastWord(s string) (size int) {
	i := len(s) - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		size++
	}
	return size
}

func main() {}
