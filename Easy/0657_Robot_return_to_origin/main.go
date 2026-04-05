// https://leetcode.com/problems/robot-return-to-origin/
package main

func judgeCircle(moves string) bool {
	var pos [2]int
	for _, ch := range moves {
		switch ch {
		case 'U':
			pos[0] += 1
		case 'D':
			pos[0] -= 1
		case 'L':
			pos[1] -= 1
		case 'R':
			pos[1] += 1
		}
	}
	return pos[0] == 0 && pos[1] == 0
}

func main() {}
