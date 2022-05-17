// https://leetcode.com/problems/spiral-matrix/
package main

func spiralOrder(matrix [][]int) []int {
	i, sx, sy, ex, ey := 0, 0, 0, len(matrix)-1, len(matrix[0])-1
	ret := make([]int, (ex+1)*(ey+1))
	for sx < ex && sy < ey {
		for y := sy; y < ey; y++ {
			ret[i] = matrix[sx][y]
			i++
		}
		for x := sx; x < ex; x++ {
			ret[i] = matrix[x][ey]
			i++
		}
		for y := ey; y > sy; y-- {
			ret[i] = matrix[ex][y]
			i++
		}
		for x := ex; x > sx; x-- {
			ret[i] = matrix[x][sy]
			i++
		}
		sx++
		sy++
		ex--
		ey--
	}
	if sx == ex && sy == ey {
		ret[i] = matrix[sx][sy]
	} else if sx == ex {
		for y := sy; y <= ey; y++ {
			ret[i] = matrix[sx][y]
			i++
		}
	} else if sy == ey {
		for x := sx; x <= ex; x++ {
			ret[i] = matrix[x][sy]
			i++
		}
	}
	return ret
}

func main() {}
