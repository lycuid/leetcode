// https://leetcode.com/problems/stone-game/
package main

func stoneGame(piles []int) bool {
	var (
		n     = len(piles)
		cache = make([][2]int, n*n)
		sum   = piles[0]
	)
	for _, pile := range piles[1:] {
		sum += pile
	}
	var solve func(int, int, int) int
	solve = func(l, r, turn int) (res int) {
		if l <= r {
			if c := cache[l*n+r]; c[0] != 0 {
				return c[1]
			}
			left, right := solve(l+1, r, 3-turn), solve(l, r-1, 3-turn)
			if turn == 1 {
				res = max(piles[l]+left, piles[r]+right)
			} else {
				res = min(left, right)
			}
			cache[l*n+r] = [2]int{turn, res}
		}
		return res
	}
	res := solve(0, n-1, 1)
	return res >= sum-res
}

func main() {}
