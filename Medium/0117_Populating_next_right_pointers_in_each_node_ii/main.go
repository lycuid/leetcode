// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	Primary, Secondary, level := 0, 1, [2][]*Node{{root}}
	if root == nil {
		goto EXIT
	}
	for ; len(level[Primary]) > 0; level[Secondary] = nil {
		var previous *Node
		for _, node := range level[Primary] {
			if previous != nil {
				previous.Next = node
			}
			previous = node
			if node.Left != nil {
				level[Secondary] = append(level[Secondary], node.Left)
			}
			if node.Right != nil {
				level[Secondary] = append(level[Secondary], node.Right)
			}
		}
		Primary, Secondary = Secondary, Primary
	}
EXIT:
	return root
}

func main() {}
