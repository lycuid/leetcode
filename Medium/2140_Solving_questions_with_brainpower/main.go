// https://leetcode.com/problems/solving-questions-with-brainpower/
package main

func mostPoints(questions [][]int) int64 {
	cache, n := make([]int64, len(questions)+1), len(questions)
	for i := n - 1; i >= 0; i-- {
		if cache[i] = int64(questions[i][0]); questions[i][1]+i+1 < n {
			cache[i] += cache[questions[i][1]+i+1]
		}
		if cache[i] < cache[i+1] {
			cache[i] = cache[i+1]
		}
	}
	return cache[0]
}

func main() {}
