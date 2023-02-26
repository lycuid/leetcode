// https://leetcode.com/problems/edit-distance/
package main

func Min(arr ...int) int {
	ret := arr[0]
	for _, num := range arr {
		if num < ret {
			ret = num
		}
	}
	return ret
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var cache [2][]int
	cache[0], cache[1] = make([]int, n+1), make([]int, n+1)
	for i := range cache[0] {
		cache[0][i] = i
	}
	for i := 1; i <= m; i++ {
		cache[1][0] = i
		for j := 1; j <= n; j++ {
			sub := cache[0][j-1]
			if word1[i-1] != word2[j-1] {
				sub++
			}
			cache[1][j] = Min(cache[0][j]+1, cache[1][j-1]+1, sub)
		}
		cache[0], cache[1] = cache[1], cache[0]
	}
	return cache[0][n]
}

func main() {}
