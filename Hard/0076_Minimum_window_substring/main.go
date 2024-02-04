// https://leetcode.com/problems/minimum-window-substring/
package main

func minWindow(s string, t string) string {
	var (
		freq       [2][123 - 65]int
		start, end int
	)

	Check := func() bool {
		for i := range freq[0] {
			if freq[0][i] > 0 && freq[1][i] < freq[0][i] {
				return false
			}
		}
		return true
	}

	for _, ch := range t {
		freq[0][ch-'A']++
	}
	for i, j := 0, 0; j < len(s); j++ {
		for freq[1][s[j]-'A']++; Check() && i <= j; i++ {
			if (start == 0 && end == 0) || j+1-i < end-start {
				start, end = i, j+1
			}
			freq[1][s[i]-'A']--
		}
	}
	return s[start:end]
}

func main() {}
