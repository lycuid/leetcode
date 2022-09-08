// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package main

// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }


func adjust(node *ListNode, number int) (int) {
  if node == nil {
    return 0
  }
  nextNumber := adjust(node.Next, number)
  if nextNumber == number {
    node.Next = node.Next.Next
  }
  return nextNumber + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  number := adjust(head, n)
  if number == n {
    return head.Next
  }
  return head
}


func main() { }

