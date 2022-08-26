// https://leetcode.com/problems/reordered-power-of-2/
package main

import "fmt"

const UPPER_BOUND int = 1e9

func MatchesDigitCount(num1, num2 int) bool {
	var freq [10]int
	for i := num1; i > 0; i /= 10 {
		freq[i%10]++
	}
	for i := num2; i > 0; i /= 10 {
		if freq[i%10]--; freq[i%10] < 0 {
			break
		}
	}
	for _, f := range freq {
		if f != 0 {
			return false
		}
	}
	return true
}

func reorderedPowerOf2(n int) bool {
	for i := 0; 1<<i < UPPER_BOUND; i++ {
		if MatchesDigitCount(n, 1<<i) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
	fmt.Println(reorderedPowerOf2(1024))
}
