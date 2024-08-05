// https://leetcode.com/problems/kth-distinct-string-in-an-array/description/
package main

func kthDistinct(arr []string, k int) string {
	repeat := make(map[string]bool)
	for _, str := range arr {
		_, found := repeat[str]
		repeat[str] = found
	}
	for _, str := range arr {
		if !repeat[str] {
			if k--; k == 0 {
				return str
			}
		}
	}
	return ""
}

func main() {}
