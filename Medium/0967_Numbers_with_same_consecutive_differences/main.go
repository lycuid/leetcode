// https://leetcode.com/problems/numbers-with-same-consecutive-differences/
package main

func numsSameConsecDiff(n int, k int) (ret []int) {
	var cache [2][]int
	Prev, Cur := 0, 1
	for i := 1; i < 10; i++ {
		cache[Cur] = append(cache[Cur], i)
	}
	for n--; n > 0; n-- {
		Prev, Cur = Cur, Prev
		for i := range cache[Prev] {
			if digit := cache[Prev][i]%10 + k; digit < 10 {
				cache[Cur] = append(cache[Cur], (cache[Prev][i]*10)+digit)
			}
			if digit := cache[Prev][i]%10 - k; k > 0 && digit >= 0 {
				cache[Cur] = append(cache[Cur], (cache[Prev][i]*10)+digit)
			}
		}
		cache[Prev] = nil
	}
	return cache[Cur]
}

func main() {}
