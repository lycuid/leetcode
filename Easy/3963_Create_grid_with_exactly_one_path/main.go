// https://leetcode.com/problems/create-grid-with-exactly-one-path/
package main

import "strings"

func createGrid(m int, n int) (res []string) {
	res = make([]string, m)
	for i := range res {
		if i == 0 {
			res[i] = strings.Repeat(".", n-1)
		} else {
			res[i] = strings.Repeat("#", n-1)
		}
		res[i] += "."
	}
	return res
}

func main() {}
