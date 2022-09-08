// https://leetcode.com/problems/peeking-iterator/
// vim:fdm=marker:fmr={{{,}}}
package main

// Not provided in problem statement. {{{
type Iterator struct {
}

func (this *Iterator) hasNext() (ret bool) {
	return
}

func (this *Iterator) next() (ret int) {
	return
}

// }}}

type PeekingIterator struct {
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
  new_iter := *this.iter
	return new_iter.next()
}

func main() {}
