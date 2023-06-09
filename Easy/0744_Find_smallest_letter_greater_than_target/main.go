// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
package main

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	if target < letters[len(letters)-1] {
		for r := len(letters) - 1; l < r; {
			if mid := (l + r) / 2; letters[mid] <= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return letters[l]
}

func main() {}
