// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
package main

func maxScore(s string) (max int) {
	zeroes := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		zeroes[i] = zeroes[i-1] + 1 - int(s[i-1]-'0')
	}
	for i, ones := len(s)-1, 0; i > 0; i-- {
		ones += int(s[i] - '0')
		if score := ones + zeroes[i]; score > max {
			max = score
		}
	}
	return max
}

func main() {}
