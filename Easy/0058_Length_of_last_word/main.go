// https://leetcode.com/problems/length-of-last-word/

package main

import "strings"

func lengthOfLastWord(s string) (l int) {
  s = strings.TrimRight(s, " ")
  i := len(s) - 1
  for i >= 0 && s[i] != ' ' {
    l++
    i--
  }
  return l
}

func main() { }

