// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) (ret *ListNode) {
	var cache []*ListNode
	for root := head; root != nil; root = root.Next {
		cache = append(cache, root)
	}
	for i := len(cache) - 1; i >= 0; i-- {
		if n--; n != 0 {
			cache[i].Next, ret = ret, cache[i]
		}
	}
	return ret
}

func main() {}
