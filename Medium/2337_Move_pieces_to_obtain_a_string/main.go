// https://leetcode.com/problems/move-pieces-to-obtain-a-string/
package main

func canChange(start string, target string) bool {
	if m, n := len(start), len(target); m == n {
		var count [2]int
		check := func(i int, fst byte) bool {
			if start[i] == fst {
				count[0]++
			}
			if target[i] == fst {
				if count[1]++; count[0] < count[1] {
					return false
				}
			}
			if snd := 'R' + 'L' - fst; start[i] == snd || target[i] == snd {
				if count[0] != count[1] {
					return false
				}
				count[0], count[1] = 0, 0
			}
			return true
		}
		for i := 0; i < n; i++ {
			if !check(i, 'R') {
				return false
			}
		}
		if count[0] != count[1] {
			return false
		}
		count[0], count[1] = 0, 0
		for i := n - 1; i >= 0; i-- {
			if !check(i, 'L') {
				return false
			}
		}
		return count[0] == count[1]
	}
	return false
}

func main() {}
