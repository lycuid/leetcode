// https://leetcode.com/problems/online-stock-span/
package main

type StockSpanner struct{ price, span []int }

func Constructor() StockSpanner { return StockSpanner{} }

func (this *StockSpanner) Next(price int) int {
	n := len(this.price)
	this.price, this.span = append(this.price, price), append(this.span, 1)
	for i := n - 1; i >= 0 && this.price[i] <= price; i -= this.span[i] {
		this.span[n] += this.span[i]
	}
	return this.span[n]
}

func main() {}
