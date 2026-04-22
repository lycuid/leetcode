// https://leetcode.com/problems/words-within-two-edits-of-dictionary/
package main

func twoEditWords(queries []string, dictionary []string) []string {
	res := make([]string, 0, len(queries))
	distance := func(a, b string) (dist int) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				dist++
			}
		}
		return dist
	}
	for _, query := range queries {
		for _, word := range dictionary {
			if distance(query, word) <= 2 {
				res = append(res, query)
				break
			}
		}
	}
	return res
}

func main() {}
