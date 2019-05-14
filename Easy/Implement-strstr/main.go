// https://leetcode.com/problems/implement-strstr/

package main

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)

	if needleLen == 0 {
		return needleLen
	}

	if haystackLen >= needleLen {
		for i := 0; i < haystackLen-needleLen+1; i++ {
			for j := 0; j < needleLen; j++ {
				if haystack[i+j] != needle[j] {
					break
				}
				if j == needleLen-1 {
					return i
				}
			}
		}
	}
	return -1
}

func main() {}
