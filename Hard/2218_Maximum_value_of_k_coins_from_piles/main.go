// https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/
package main

type Solution struct{ cache [][]int }

func MakeSolution(n, k int) Solution {
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, k)
	}
	return Solution{cache}
}

func (this *Solution) Solve(piles [][]int, i, k int) (ret int) {
	if i >= len(piles) || k == 0 {
		return 0
	}
	if this.cache[i][k] > 0 {
		return this.cache[i][k]
	}
	ret = this.Solve(piles, i+1, k)
	for current, j := 0, 0; j < len(piles[i]) && j < k; j++ {
		current += piles[i][j]
		if n := this.Solve(piles, i+1, k-j-1) + current; ret < n {
			ret = n
		}
	}
	this.cache[i][k] = ret
	return ret
}

func maxValueOfCoins(piles [][]int, k int) int {
	solution := MakeSolution(len(piles)+1, k+1)
	return solution.Solve(piles, 0, k)
}

func main() {}
