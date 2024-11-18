// https://leetcode.com/problems/defuse-the-bomb/
package main

func decrypt(code []int, k int) []int {
	ret, n := make([]int, len(code)), len(code)
	if inc := 1; k != 0 {
		c := func(i int) int { return (i + n) % n }
		if k < 0 {
			inc = -1
		}
		for i := 0; i != k; i += inc {
			ret[0] += code[c(i)]
		}
		ret[0] += code[c(k)] - code[0]
		for i, index := 1, c(inc); i < len(ret); i, index = i+1, c(index+inc) {
			ret[index] = ret[c(index-inc)] + code[c(index+k)] - code[index]
		}
	}
	return ret
}

func main() {}
