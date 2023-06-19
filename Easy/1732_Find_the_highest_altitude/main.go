// https://leetcode.com/problems/find-the-highest-altitude/
package main

func largestAltitude(gain []int) (ret int) {
	for i, altitude := 0, 0; i < len(gain); i++ {
		if altitude += gain[i]; altitude > ret {
			ret = altitude
		}
	}
	return ret
}

func main() {}
