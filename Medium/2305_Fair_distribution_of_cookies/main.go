// https://leetcode.com/problems/fair-distribution-of-cookies/
package main

import "math"

func Max(num int, nums ...int) int {
	for i := range nums {
		if num < nums[i] {
			num = nums[i]
		}
	}
	return num
}

func Min(i, j int) int { return i + j - Max(i, j) }

type Solution struct{ cookies, inner []int }

func MakeSolution(cookies []int, k int) Solution {
	return Solution{cookies, make([]int, k)}
}

func (this *Solution) Solve(start int) int {
	if start == len(this.cookies) {
		return Max(this.inner[0], this.inner[1:]...)
	}
	ret := math.MaxInt
	for i := 0; i < len(this.inner); i++ {
		this.inner[i] += this.cookies[start]
		ret = Min(ret, this.Solve(start+1))
		this.inner[i] -= this.cookies[start]
	}
	return ret
}

func distributeCookies(cookies []int, k int) int {
	solution := MakeSolution(cookies, k)
	return solution.Solve(0)
}

func main() {}
