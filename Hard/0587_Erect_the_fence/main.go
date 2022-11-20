// https://leetcode.com/problems/erect-the-fence/
package main

import "sort"

func Cross(p1, p2, p3 []int) int {
	return ((p3[0] - p2[0]) * (p2[1] - p1[1])) - ((p3[1] - p2[1]) * (p2[0] - p1[0]))
}

func outerTrees(trees [][]int) (ret [][]int) {
	var onStack [101][101]bool
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] == trees[j][0] {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})
	for _, tree := range trees {
		for n := len(ret); n >= 2 && Cross(ret[n-2], ret[n-1], tree) > 0; n = len(ret) {
			ret, onStack[ret[n-1][0]][ret[n-1][1]] = ret[:n-1], false
		}
		if !onStack[tree[0]][tree[1]] {
			ret, onStack[tree[0]][tree[1]] = append(ret, tree), true
		}
	}
	for i := len(trees) - 1; i >= 0; i-- {
		tree := trees[i]
		for n := len(ret); n >= 2 && Cross(ret[n-2], ret[n-1], tree) > 0; n = len(ret) {
			ret, onStack[ret[n-1][0]][ret[n-1][1]] = ret[:n-1], false
		}
		if !onStack[tree[0]][tree[1]] {
			ret, onStack[tree[0]][tree[1]] = append(ret, tree), true
		}
	}
	return ret
}

func main() {}
