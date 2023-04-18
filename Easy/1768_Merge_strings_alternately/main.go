// https://leetcode.com/problems/merge-strings-alternately/
package main

func mergeAlternately(w1 string, w2 string) string {
	var ret []byte
	for ; len(w1) > 0 && len(w2) > 0; w1, w2 = w1[1:], w2[1:] {
		ret = append(ret, w1[0], w2[0])
	}
	return string(ret) + w1 + w2
}

func main() {}
