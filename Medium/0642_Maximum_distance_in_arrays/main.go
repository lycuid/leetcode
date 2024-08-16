// https://leetcode.com/problems/maximum-distance-in-arrays/
package main

func maxDistance(arrays [][]int) (dist int) {
	Abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	l, r := arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		fst, snd := arrays[i][0], arrays[i][len(arrays[i])-1]
		dist = max(dist, max(Abs(r-fst), Abs(snd-l)))
		l, r = min(l, fst), max(r, snd)
	}
	return dist
}

func main() {}
