// https://leetcode.com/problems/longest-string-chain/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Distance(ss, sb int, small, big string, distance int) bool {
	if distance >= 0 {
		for ; ss < len(small); ss, sb = ss+1, sb+1 {
			if small[ss] != big[sb] {
				return Distance(ss, sb+1, small, big, distance-1)
			}
		}
	}
	return sb+distance == len(big)
}

func longestStrChain(words []string) (ret int) {
	var levels [17][]string
	var dp [17][]int
	for _, word := range words {
		i := len(word)
		levels[i], dp[i] = append(levels[i], word), append(dp[i], 1)
	}
	for i := 1; i < 16; i++ {
		for j, from := range levels[i] {
			for k, to := range levels[i+1] {
				if Distance(0, 0, from, to, 1) {
					dp[i+1][k] = Max(dp[i+1][k], dp[i][j]+1)
				}
				ret = Max(ret, dp[i+1][k])
			}
		}
	}
	return Max(1, ret)
}

func main() {}
