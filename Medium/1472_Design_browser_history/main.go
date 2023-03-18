// https://leetcode.com/problems/design-browser-history/
package main

type Node struct {
	Value      string
	Prev, Next *Node
}

type BrowserHistory struct{ current *Node }

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{&Node{Value: homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	node := Node{Value: url, Prev: this.current}
	this.current.Next, this.current = &node, &node
}

func (this *BrowserHistory) Back(steps int) string {
	for ; this.current.Prev != nil && steps > 0; steps-- {
		this.current = this.current.Prev
	}
	return this.current.Value
}

func (this *BrowserHistory) Forward(steps int) string {
	for ; this.current.Next != nil && steps > 0; steps-- {
		this.current = this.current.Next
	}
	return this.current.Value
}

func main() {}
