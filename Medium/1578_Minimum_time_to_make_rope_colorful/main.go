// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package main

func minCost(cs string, time []int) (ret int) {
	for i, j, max, sum := 0, 0, 0, 0; i < len(cs); i++ {
		max, sum = time[i], time[i]
		for j = i + 1; j < len(cs) && cs[i] == cs[j]; i, j = i+1, j+1 {
			if sum += time[j]; time[j] > max {
				max = time[j]
			}
		}
		if sum != max { // more than 1 repeat (1 <= time[i] <= 1e4)
			ret += sum - max
		}
	}
	return ret
}

func main() {}
