// https://leetcode.com/problems/count-and-say/

package main

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	var (
		result bytes.Buffer
		i      int
	)

	previous := countAndSay(n - 1)
	counter := 1

	for i = 1; i < len(previous); i++ {
		if previous[i] != previous[i-1] {
			result.WriteString(strconv.Itoa(counter))
			result.WriteString(string(previous[i-1]))
			counter = 1
			continue
		}
		counter++
	}
	result.WriteString(strconv.Itoa(counter))
	result.WriteString(string(previous[i-1]))

	return result.String()
}

func main() {}
