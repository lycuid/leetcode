// https://leetcode.com/problems/make-the-string-great/
package main

func makeGood(s string) string {
	var q []byte
	for i := range s {
		if n := len(q); n > 0 && (q[n-1] == s[i]-32 || q[n-1]-32 == s[i]) {
			q = q[:n-1]
		} else {
			q = append(q, s[i])
		}
	}
	return string(q)
}

func main() {}
