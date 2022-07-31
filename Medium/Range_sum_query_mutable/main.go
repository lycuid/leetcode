// https://leetcode.com/problems/range-sum-query-mutable/
package main

type NumArray struct {
	inner []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) Update(index int, val int) {
	this.inner[index] = val
}

func (this *NumArray) SumRange(left int, right int) (ret int) {
	for i := left; i <= right; i++ {
		ret += this.inner[i]
	}
	return ret
}

func main() {}
