// https://leetcode.com/problems/walking-robot-simulation/
package main

import "sort"

const (
	N = iota
	E = iota
	S = iota
	W = iota
)

func robotSim(commands []int, obstacles [][]int) (ret int) {
	x, y := make(map[int][]int), make(map[int][]int)
	for _, c := range obstacles {
		x[c[0]] = append(x[c[0]], c[1])
		y[c[1]] = append(y[c[1]], c[0])
	}
	for key := range x {
		sort.Ints(x[key])
	}
	for key := range y {
		sort.Ints(y[key])
	}
	var dx, dy, dir int
	for _, cmd := range commands {
		switch cmd {
		case -1:
			dir = (dir + 1) % 4
		case -2:
			dir = (dir + 3) % 4
		default:
			switch dir {
			case N:
				col := x[dx]
				i := sort.Search(len(col), func(i int) bool {
					return col[i] > dy
				})
				if i < len(col) && col[i] <= dy+cmd {
					dy = col[i] - 1
				} else {
					dy += cmd
				}
			case E:
				row := y[dy]
				i := sort.Search(len(row), func(i int) bool {
					return row[i] > dx
				})
				if i != len(row) && row[i] <= dx+cmd {
					dx = row[i] - 1
				} else {
					dx += cmd
				}
			case S:
				col := x[dx]
				i := sort.Search(len(col), func(i int) bool {
					return col[i] >= dy
				}) - 1
				if i >= 0 && col[i] >= dy-cmd {
					dy = col[i] + 1
				} else {
					dy -= cmd
				}
			case W:
				row := y[dy]
				i := sort.Search(len(row), func(i int) bool {
					return row[i] >= dx
				}) - 1
				if i >= 0 && row[i] >= dx-cmd {
					dx = row[i] + 1
				} else {
					dx -= cmd
				}
			}
			ret = max(ret, dx*dx+dy*dy)
		}
	}
	return ret
}

func main() {}
