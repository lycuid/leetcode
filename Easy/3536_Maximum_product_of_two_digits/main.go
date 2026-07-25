// https://leetcode.com/problems/maximum-product-of-two-digits/
package main

func maxProduct(n int) int {
	fst, snd := n%10, (n/10)%10
	if fst < snd {
		fst, snd = snd, fst
	}
	for n /= 100; n > 0; n /= 10 {
		if d := n % 10; d >= fst {
			fst, snd = d, fst
		} else if d > snd {
			snd = d
		}
	}
	return fst * snd
}

func main() {}
