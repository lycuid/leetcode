// https://leetcode.com/problems/count-of-smaller-numbers-after-self/
package main

type Tuple struct {
	index, val int
}

func Merge(start, mid, end int, arr []Tuple, ret []int) {
	var tmp []Tuple
	i1, j1, i2, j2 := start, mid, mid+1, end
	for i1 <= j1 && i2 <= j2 {
		if arr[i1].val <= arr[i2].val {
			tmp, i2 = append(tmp, arr[i2]), i2+1
		} else {
			ret[arr[i1].index] += (j2 - i2 + 1)
			tmp, i1 = append(tmp, arr[i1]), i1+1
		}
	}
	for i1 <= j1 {
		tmp, i1 = append(tmp, arr[i1]), i1+1
	}
	for i2 <= j2 {
		tmp, i2 = append(tmp, arr[i2]), i2+1
	}
	for i := range tmp {
		arr[start+i] = tmp[i]
	}
}

func Sort(start, end int, arr []Tuple, ret []int) {
	if start < end {
		mid := (start + end) / 2
		Sort(start, mid, arr, ret)
		Sort(mid+1, end, arr, ret)
		Merge(start, mid, end, arr, ret)
	}
}

func countSmaller(nums []int) []int {
	arr, ret := make([]Tuple, len(nums)), make([]int, len(nums))
	for i, num := range nums {
		arr[i] = Tuple{i, num}
	}
	Sort(0, len(nums)-1, arr, ret)
	return ret
}

func main() {}
