// https://leetcode.com/problems/push-dominoes/
package main

func pushDominoes(dominoes string) string {
	xs := []byte(dominoes)
	L, R := make([]int, len(xs)), make([]int, len(xs))
	for i := 0; i < len(xs); i++ {
		if xs[i] == 'R' {
			for R[i], i = 0, i+1; i < len(xs) && xs[i] == '.'; i++ {
				R[i] = R[i-1] + 1
			}
			i--
		}
	}
	for i := len(xs) - 1; i >= 0; i-- {
		if xs[i] == 'L' {
			for L[i], i = 0, i-1; i >= 0 && xs[i] == '.'; i-- {
				L[i] = L[i+1] + 1
			}
			i++
		}
	}
	for i := range xs {
		if L[i] > 0 && R[i] > 0 {
			if L[i] < R[i] {
				xs[i] = 'L'
			} else if L[i] > R[i] {
				xs[i] = 'R'
			}
		} else if L[i] > 0 {
			xs[i] = 'L'
		} else if R[i] > 0 {
			xs[i] = 'R'
		}
	}
	return string(xs)
}

func main() {}
