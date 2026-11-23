// https://leetcode.com/problems/product-of-the-last-k-numbers/
package main

type ProductOfNumbers struct{ items []int }

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.items = this.items[:1]
	} else {
		this.items = append(this.items, this.items[len(this.items)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if n := len(this.items) - 1; n >= k {
		return this.items[n] / this.items[n-k]
	}
	return 0

}

func main() {}
