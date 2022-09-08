// https://leetcode.com/problems/poor-pigs/
package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) /
		math.Log(float64(minutesToTest/minutesToDie+1))))
}

func main() {}
