// https://leetcode.com/problems/weighted-word-mapping/
package main

func mapWordWeights(words []string, weights []int) string {
	res := make([]byte, len(words))
	for i, word := range words {
		var weight int
		for _, ch := range word {
			weight += weights[int(ch-'a')]
		}
		res[i] = byte('z' - weight%26)
	}
	return string(res)
}

func main() {}
