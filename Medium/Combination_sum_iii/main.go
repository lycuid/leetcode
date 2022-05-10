// https://leetcode.com/problems/combination-sum-iii/
package main

func Aux(start, end, count, target int) (masks []int) {
	if count == 1 {
		if target == start {
			masks = []int{1 << target}
		}
		goto EXIT
	}
	for i := start + 1; i <= end; i++ {
		aux := Aux(i, end, count-1, target-start)
		for j := range aux {
			aux[j] |= 1 << start
		}
		masks = append(masks, aux...)
	}
EXIT:
	return masks
}

func combinationSum3(k, n int) (ret [][]int) {
	for i := 1; i <= 9; i++ {
		for _, mask := range Aux(i, 9, k, n) {
			if mask == 0 {
				continue
			}
			var tmp []int
			for j := 1; j <= 9; j++ {
				if mask&(1<<j) > 0 {
					tmp = append(tmp, j)
				}
			}
			ret = append(ret, tmp)
		}
	}
	return ret
}

func main() {}
