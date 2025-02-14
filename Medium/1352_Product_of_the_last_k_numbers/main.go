// https://leetcode.com/problems/product-of-the-last-k-numbers/
package main

type ProductOfNumbers struct{ nums []int }

func Constructor() ProductOfNumbers { return ProductOfNumbers{[]int{1}} }

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.nums = this.nums[:1]
	} else {
		this.nums = append(this.nums, this.nums[len(this.nums)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if n := len(this.nums) - 1; n >= k {
		return this.nums[n] / this.nums[n-k]
	}
	return 0
}

func main() {}
