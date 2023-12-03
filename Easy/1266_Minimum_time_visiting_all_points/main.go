// https://leetcode.com/problems/minimum-time-visiting-all-points/
package main

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func minTimeToVisitAllPoints(ps [][]int) (count int) {
	for i := 0; i < len(ps)-1; i++ {
		if n, m := Abs(ps[i+1][0]-ps[i][0]), Abs(ps[i+1][1]-ps[i][1]); n > m {
			count += n
		} else {
			count += m
		}
	}
	return count
}

func main() {}
