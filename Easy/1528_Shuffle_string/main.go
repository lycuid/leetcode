// https://leetcode.com/problems/shuffle-string/
package main

func restoreString(s string, indices []int) string {
	ret := make([]byte, len(indices))
	for i := range indices {
		ret[indices[i]] = s[i]
	}
	return string(ret)
}

func main() {}
