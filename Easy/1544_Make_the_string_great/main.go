// https://leetcode.com/problems/make-the-string-great/
package main

func makeGood(s string) string {
	ret, i := []byte(s), 0
	Abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for j := 0; j < len(s); j++ {
		if i > 0 && Abs(int(ret[j])-int(ret[i-1])) == 32 {
			i--
		} else {
			ret[i], i = ret[j], i+1
		}
	}
	return string(ret[:i])
}

func main() {}
