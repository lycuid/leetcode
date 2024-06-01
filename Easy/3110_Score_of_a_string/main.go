// https://leetcode.com/problems/score-of-a-string/
package main

func scoreOfString(s string) (score int) {
	for i, val := 1, 0; i < len(s); i++ {
		if val = int(s[i-1]) - int(s[i]); val < 0 {
			val = -val
		}
		score += val
	}
	return score
}

func main() {}
