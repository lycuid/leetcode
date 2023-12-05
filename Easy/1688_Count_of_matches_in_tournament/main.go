// https://leetcode.com/problems/count-of-matches-in-tournament/
package main

func numberOfMatches(n int) (count int) {
	for ; n > 1; n = (n / 2) + (n % 2) {
		count += n / 2
	}
	return count
}

func main() {}
