// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
package main

func countNegatives(grid [][]int) (count int) {
	for _, row := range grid {
		l, r := 0, len(row)
		for l < r {
			if mid := l + (r-l)/2; row[mid] < 0 {
				r = mid
			} else {
				l = mid + 1
			}
		}
		count += len(row) - l
	}
	return count
}

func main() {}
