// https://leetcode.com/problems/merge-k-sorted-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type PriorityQueue struct{ inner []*ListNode }

func MakePQ() PriorityQueue            { return PriorityQueue{make([]*ListNode, 1)} }
func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }

func (this *PriorityQueue) Push(item *ListNode) {
	if item != nil {
		this.inner = append(this.inner, item)
		for i := len(this.inner) - 1; i > 1 && this.inner[i].Val < this.inner[i/2].Val; i /= 2 {
			this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
		}
	}
}

func (this *PriorityQueue) Pop() *ListNode {
	if this.Empty() {
		return nil
	}
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.inner[j+1].Val < this.inner[j].Val {
			j++
		}
		if this.inner[i].Val < this.inner[j].Val {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func mergeKLists(lists []*ListNode) (head *ListNode) {
	pq := MakePQ()
	for i := range lists {
		pq.Push(lists[i])
	}
	if head = pq.Pop(); head != nil {
		pq.Push(head.Next)
	}
	for last := head; !pq.Empty(); last = last.Next {
		last.Next = pq.Pop()
		pq.Push(last.Next.Next)
	}
	return head
}

func main() {}
