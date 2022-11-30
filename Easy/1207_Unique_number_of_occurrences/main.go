// https://leetcode.com/problems/unique-number-of-occurrences/
package main

func uniqueOccurrences(nums []int) bool {
	freq, found := [2001]int{}, [2001]bool{}
	for _, num := range nums {
		freq[num+1000]++
	}
	for i := range freq {
		if freq[i] > 0 && found[freq[i]] {
			return false
		}
		found[freq[i]] = true
	}
	return true
}

func main() {}
