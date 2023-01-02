// https://leetcode.com/problems/detect-capital/
package main

func detectCapitalUse(word string) (all_cap bool) {
	is_cap := func(i int) bool { return word[i] >= 'A' && word[i] <= 'Z' }
	if len(word) > 1 && is_cap(0) {
		all_cap = is_cap(1)
	}
	for i := 1; i < len(word); i++ {
		if is_cap(i) != all_cap {
			return false
		}
	}
	return true
}

func main() {}
