// https://leetcode.com/problems/copy-list-with-random-pointer/
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) (ret *Node) {
	if head != nil {
		for root := head; root != nil; root = root.Next.Next {
			node := &Node{Val: root.Val, Next: root.Next, Random: root.Random}
			root.Next = node
		}
		for root := head; root != nil; root = root.Next.Next {
			if clone := root.Next; clone.Random != nil {
				clone.Random = clone.Random.Next
			}
		}
		ret = head.Next
		for root := head; root != nil; root = root.Next {
			clone := root.Next
			if root.Next = clone.Next; clone.Next != nil {
				clone.Next = clone.Next.Next
			}
		}
	}
	return ret
}

func main() {}
