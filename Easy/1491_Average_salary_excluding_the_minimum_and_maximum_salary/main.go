// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
package main

import "math"

func average(salary []int) (ret float64) {
	min, max := math.MaxFloat64, 0.
	for i := range salary {
		num := float64(salary[i])
		ret, min, max = ret+num, math.Min(min, num), math.Max(max, num)
	}
	return (ret - min - max) / float64(len(salary)-2)
}

func main() {}
