// https://leetcode.com/problems/find-the-highest-altitude/
package main

func largestAltitude(gain []int) (res int) {
	for i, alt := 0, 0; i < len(gain); i++ {
		alt += gain[i]
		res = max(res, alt)
	}
	return res
}

func main() {}
