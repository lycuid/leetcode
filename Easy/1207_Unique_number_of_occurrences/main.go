// https://leetcode.com/problems/unique-number-of-occurrences/
package main

func uniqueOccurrences(arr []int) bool {
	freq, found := make(map[int]int), make(map[int]bool)
	for i := range arr {
		freq[arr[i]]++
	}
	for i := range freq {
		if found[freq[i]] {
			return false
		}
		found[freq[i]] = true
	}
	return true
}

func main() {}
