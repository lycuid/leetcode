// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
package main

func minCost(colors string, neededTime []int) (cost int) {
	for i, j := 0, 1; j < len(colors); i, j = i+1, j+1 {
		if colors[i] == colors[j] {
			if neededTime[i] > neededTime[j] {
				neededTime[i], neededTime[j] = neededTime[j], neededTime[i]
			}
			cost += neededTime[i]
		}
	}
	return cost
}

func main() {}
