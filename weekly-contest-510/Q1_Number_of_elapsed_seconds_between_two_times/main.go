// https://leetcode.com/contest/weekly-contest-510/problems/number-of-elapsed-seconds-between-two-times/
package main

import "time"

func secondsBetweenTimes(startTime string, endTime string) int {
	start, _ := time.Parse("15:04:05", startTime)
	end, _ := time.Parse("15:04:05", endTime)
	return int(end.Sub(start).Seconds())
}

func main() {}
