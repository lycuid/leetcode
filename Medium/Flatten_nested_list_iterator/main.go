// https://leetcode.com/problems/flatten-nested-list-iterator/
package main

type NestedIterator struct {
	iter   []int
	cursor int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var iter []int
	flatten(nestedList, &iter)
	return &NestedIterator{iter: iter}
}

func (this *NestedIterator) Next() (ret int) {
	ret = this.iter[this.cursor]
	this.cursor++
	return
}

func (this *NestedIterator) HasNext() bool {
	return this.cursor < len(this.iter)
}

func flatten(nested []*NestedInteger, iter *[]int) {
	for _, input := range nested {
		if input.IsInteger() {
			*iter = append(*iter, input.GetInteger())
		} else {
			flatten(input.GetList(), iter)
		}
	}
}

func main() {}

// =============================================================================
// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() (ret bool) {
	return
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() (ret int) {
	return
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() (ret []*NestedInteger) {
	return
}
