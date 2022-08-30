// https://leetcode.com/problems/rotate-image/
package main

func rotate(m [][]int) {
	n := len(m) - 1
	for i := 0; i <= n/2; i++ {
		for j := i; j < n-i; j++ {
			m[i][j], m[j][n-i], m[n-i][n-j], m[n-j][i] =
				m[n-j][i], m[i][j], m[j][n-i], m[n-i][n-j]
		}
	}
}
