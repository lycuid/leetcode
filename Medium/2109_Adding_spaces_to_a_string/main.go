// https://leetcode.com/problems/adding-spaces-to-a-string/
package main

import "strings"

func addSpaces(s string, spaces []int) string {
	var builder strings.Builder
	for i, ch := range s {
		if len(spaces) > 0 && spaces[0] == i {
			builder.WriteByte(' ')
			spaces = spaces[1:]
		}
		builder.WriteRune(ch)
	}
	return builder.String()
}

func main() {}
