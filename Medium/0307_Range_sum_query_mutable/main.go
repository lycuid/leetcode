// https://leetcode.com/problems/range-sum-query-mutable/
package main

type NumArray struct{ tree []int }

func Constructor(nums []int) NumArray {
	n := len(nums)
	for n&(n-1) != 0 {
		n++
	}
	tree := make([]int, 2*n)
	for i := range nums {
		tree[n+i] = nums[i]
	}
	for i := n - 1; i >= 1; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	return NumArray{tree}
}

func (this *NumArray) Update(index int, val int) {
	index += (len(this.tree) / 2)
	this.tree[index] = val
	for index /= 2; index >= 1; index /= 2 {
		this.tree[index] = this.tree[index*2] + this.tree[index*2+1]
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	var q func(int, int, int) int
	q = func(node, low, high int) int {
		if low >= left && high <= right {
			return this.tree[node]
		}
		if low > right || high < left {
			return 0
		}
		mid := (low + high) / 2
		return q(node*2, low, mid) + q(node*2+1, mid+1, high)
	}
	return q(1, 0, len(this.tree)/2-1)
}

func main() {}
