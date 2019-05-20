// https://leetcode.com/problems/zigzag-conversion/

package main

import (
	"bytes"
)

func getDistances(n int) []int {
	distances := make([]int, n)
	distances[n-1] = 2*n - 2

	for i := 0; i < n-1; i++ {
		distances[i] = 2*n - (2 * (i + 1))
	}

	return distances
}

func convert(s string, n int) string {
	if n <= 1 {
		return s
	}

	var zigZag bytes.Buffer

	distances := getDistances(n)
	dlen := len(distances)

	for i := 0; i < dlen; i++ {
		pointer := i
		for j := 0; pointer < len(s); j++ {
			zigZag.WriteByte(s[pointer])
			if j%2 == 0 {
				pointer += distances[i]
			} else {
				pointer += distances[dlen-i-1]
			}
		}
	}

	return zigZag.String()
}

func main() {}
