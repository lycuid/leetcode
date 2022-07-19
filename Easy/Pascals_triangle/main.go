// https://leetcode.com/problems/pascals-triangle/submissions/
package main

func generate(numRows int) (ret [][]int) {
	ret = make([][]int, numRows)
	for i := range ret {
		ret[i] = make([]int, i+1)
	}
	for i := range ret {
		ret[i][0], ret[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret
}

func main() {}
