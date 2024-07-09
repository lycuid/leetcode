// https://leetcode.com/problems/average-waiting-time/
package main

func averageWaitingTime(customers [][]int) float64 {
	total, waiting := customers[0][1], customers[0][1]
	for i := 1; i < len(customers); i++ {
		total = max(0, total-customers[i][0]+customers[i-1][0]) + customers[i][1]
		waiting += total
	}
	return float64(waiting) / float64(len(customers))
}

func main() {}
