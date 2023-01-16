// https://leetcode.com/problems/insert-interval/
package main

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{newInterval}
	for i := 0; i < len(intervals); i++ {
		fst, snd := ret[len(ret)-1], intervals[i]
		if snd[0] < fst[0] {
			fst, snd = snd, fst
		}
		if ret[len(ret)-1] = fst; fst[1] >= snd[0] {
			if fst[1] < snd[1] {
				ret[len(ret)-1][1] = snd[1]
			}
		} else {
			ret = append(ret, snd)
		}
	}
	return ret
}

func main() {}
