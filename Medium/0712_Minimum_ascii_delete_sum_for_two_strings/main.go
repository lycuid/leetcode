// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
package main

func minimumDeleteSum(s1 string, s2 string) int {
	curr, prev := make([]int, len(s2)+1), make([]int, len(s2)+1)
	for i := range s1 {
		curr, prev = prev, curr
		for j := range s2 {
			curr[j+1] = max(curr[j], prev[j+1])
			if s1[i] == s2[j] {
				curr[j+1] = max(curr[j+1], int(s1[i])+prev[j])
			}
		}
	}
	sum := func(slice []byte) (total int) {
		for _, ch := range slice {
			total += int(ch)
		}
		return total
	}
	return sum([]byte(s1)) + sum([]byte(s2)) - curr[len(s2)]*2
}

func main() {}
