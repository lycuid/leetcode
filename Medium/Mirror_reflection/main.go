// https://leetcode.com/problems/mirror-reflection/
package main

func gcd(x, y int) int {
	var smol, big int
	if smol, big = x, y; smol > big {
		smol, big = big, smol
	}
	if big%smol == 0 {
		return smol
	}
	return gcd(big-smol, smol)
}

func mirrorReflection(p int, q int) (ret int) {
	if n := (p * q) / gcd(p, q); (n/p)%2 == 1 {
		if ret++; (n/q)%2 == 0 {
			ret++
		}
	}
	return ret
}

func main() {}
