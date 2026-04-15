// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/
package main

func closestTarget(words []string, target string, startIndex int) int {
	res := len(words)
	for i, word := range words {
		if index := i; word == target {
			if index < startIndex {
				index += len(words)
			}
			res = min(res, index-startIndex, len(words)+startIndex-index)
		}
	}
	if res >= len(words) {
		return -1
	}
	return res
}

func main() {}
