// https://leetcode.com/problems/count-sorted-vowel-strings/
package main

func countVowelStrings(n int) int {
	a := []int{1, 1, 1, 1, 1}
	for ; n > 0; n-- {
		for i := 3; i >= 0; i-- {
			a[i] += a[i+1]
		}
	}
	return a[0]
}

func main() {}
