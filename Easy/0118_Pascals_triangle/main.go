// https://leetcode.com/problems/pascals-triangle/
package main

func generate(numRows int) [][]int {
	ret := make([][]int, 0, numRows)
	ret = append(ret, []int{1})
	for i := 1; i < numRows; i++ {
		tmp := make([]int, 0, len(ret[i-1])+2)
		tmp = append(tmp, 1)
		for j := 0; j < len(ret[i-1])-1; j++ {
			tmp = append(tmp, ret[i-1][j]+ret[i-1][j+1])
		}
		tmp = append(tmp, 1)
		ret = append(ret, tmp)
	}
	return ret
}

func main() {}
