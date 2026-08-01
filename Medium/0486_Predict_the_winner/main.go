// https://leetcode.com/problems/predict-the-winner/
package main

func predictTheWinner(nums []int) bool {
	var (
		n     = len(nums)
		sum   = nums[0]
		cache = make([][][2]int, n)
	)
	for _, num := range nums[1:] {
		sum += num
	}
	for i := range cache {
		cache[i] = make([][2]int, n)
	}
	var solve func(start, end int, turn int) int
	solve = func(l, r int, turn int) (res int) {
		if l <= r {
			if c := cache[l][r]; c[0] != 0 {
				return c[1]
			}
			left, right := solve(l+1, r, 3-turn), solve(l, r-1, 3-turn)
			if turn == 1 {
				res = max(nums[l]+left, nums[r]+right)
			} else {
				res = min(left, right)
			}
			cache[l][r] = [2]int{turn, res}
		}
		return res
	}
	res := solve(0, n-1, 1)
	return res >= sum-res
}

func main() {}
