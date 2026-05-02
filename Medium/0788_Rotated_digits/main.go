// https://leetcode.com/problems/rotated-digits/
package main

func rotatedDigits(n int) (count int) {
outer:
	for i := 1; i <= n; i++ {
		var found bool
		for j := i; j > 0; j /= 10 {
			switch d := j % 10; d {
			case 3, 4, 7:
				continue outer
			case 2, 5, 6, 9:
				found = true
			}
		}
		if found {
			count++
		}
	}
	return count
}

func main() {}
