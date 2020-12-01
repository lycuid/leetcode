// https://leetcode.com/problems/merge-k-sorted-lists/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		head     *ListNode
		previous *ListNode
	)
	l := len(lists)

	for {
		smallest := -1
		for i := 0; i < l; i++ {
			current := lists[i]
			if current == nil {
				continue
			}
			if smallest == -1 || (*current).Val < lists[smallest].Val {
				smallest = i
			}
		}
		if smallest == -1 { break; }
		if previous == nil {
			head = lists[smallest]
			previous = lists[smallest]
		} else {
			previous.Next = lists[smallest]
			previous = lists[smallest]
		}
		lists[smallest] = (*lists[smallest]).Next
	}

	return head
}

func main() { }

