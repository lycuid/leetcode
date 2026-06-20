// https://leetcode.com/problems/maximum-building-height/
package main

import (
	"math"
	"sort"
)

func maxBuilding(n int, restrictions [][]int) (res int) {
	restrictions = append(restrictions, []int{1, 0}, []int{n, math.MaxInt})
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})
	for i := 0; i < len(restrictions)-1; i++ {
		restrictions[i+1][1] = min(restrictions[i+1][1], restrictions[i][1]+restrictions[i+1][0]-restrictions[i][0])
	}
	for i := len(restrictions) - 1; i > 0; i-- {
		restrictions[i-1][1] = min(restrictions[i-1][1], restrictions[i][1]+restrictions[i][0]-restrictions[i-1][0])
		ri, rj := restrictions[i-1], restrictions[i]
		t := (rj[0] - ri[0] - 1) + ri[1]
		res = max(res, t-(t-rj[1])>>1)
	}
	return res
}

func main() {}
