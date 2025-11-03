// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package main

func minCost(colors string, neededTime []int) (cost int) {
	for i := 0; i < len(colors)-1; i++ {
		if colors[i] == colors[i+1] {
			if neededTime[i] > neededTime[i+1] {
				neededTime[i], neededTime[i+1] = neededTime[i+1], neededTime[i]
			}
			cost += neededTime[i]
		}
	}
	return cost
}

func main() {}
