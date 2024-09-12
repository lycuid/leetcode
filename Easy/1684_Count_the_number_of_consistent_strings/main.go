// https://leetcode.com/problems/count-the-number-of-consistent-strings/
package main

func countConsistentStrings(allowed string, words []string) (count int) {
	var num int64
	for i := 0; i < len(allowed); i++ {
		num |= 1 << int64(allowed[i]-'a')
	}
mainloop:
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if num>>(word[i]-'a')&1 == 0 {
				continue mainloop
			}
		}
		count++
	}
	return count
}

func main() {}
