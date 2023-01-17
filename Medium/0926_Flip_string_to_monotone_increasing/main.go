// https://leetcode.com/problems/flip-string-to-monotone-increasing/
package main

func minFlipsMonoIncr(s string) int {
	var ret [2]int
	for _, ch := range s {
		if ret[ch-'0']++; ret[1] < ret[0] {
			ret[0] = ret[1]
		}
	}
	return ret[0]
}

func main() {}
