// https://leetcode.com/problems/maximum-product-of-word-lengths/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProduct(words []string) (count int) {
	l := len(words)
	bs := make([]uint32, l)
	for i := range words {
		for _, ch := range words[i] {
			bs[i] |= (1 << (ch - 'a'))
		}
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if bs[i]&bs[j] == 0 {
				count = Max(count, len(words[i])*len(words[j]))
			}
		}
	}
	return count
}

func main() {}
