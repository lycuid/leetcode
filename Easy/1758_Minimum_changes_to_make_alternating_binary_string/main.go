// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
package main

func minOperations(s string) int {
	var count [2]int
	for i, curr := 0, byte('1'); i < len(s); i, curr = i+1, '1'+'0'-curr {
		if s[i] == curr {
			count[0]++
		} else {
			count[1]++
		}
	}
	if count[0] < count[1] {
		return count[0]
	}
	return count[1]
}

func main() {}
