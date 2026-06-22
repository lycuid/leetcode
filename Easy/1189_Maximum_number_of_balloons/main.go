// https://leetcode.com/problems/maximum-number-of-balloons/
package main

func maxNumberOfBalloons(text string) int {
	var b, a, l, o, n int
	for _, ch := range text {
		switch ch {
		case 'b':
			b++
		case 'a':
			a++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}
	return min(b, a, l>>1, o>>1, n)
}

func main() {}
