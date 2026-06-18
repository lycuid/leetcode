// https://leetcode.com/problems/angle-between-hands-of-a-clock/
package main

func angleClock(hour int, minutes int) float64 {
	h, m := float64(hour%12)*30, float64(minutes)*6
	if h += 30 * (m / 360); h > m {
		h, m = m, h
	}
	return min(360-m+h, m-h)
}

func main() {}
