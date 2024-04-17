// https://leetcode.com/problems/smallest-string-starting-from-leaf/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Smaller(s1, s2 string) string {
	var i int
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] < s2[i] {
			return s1
		}
		if s1[i] > s2[i] {
			return s2
		}
	}
	if i == len(s1) {
		return s1
	}
	return s2
}

func Reverse(str string) string {
	xs := []byte(str)
	for i, n := 0, len(xs); i < n/2; i++ {
		xs[i], xs[n-i-1] = xs[n-i-1], xs[i]
	}
	return string(xs)
}

func Aux(root *TreeNode, path string) string {
	path += string(byte(root.Val) + 'a')
	if root.Left != nil || root.Right != nil {
		if root.Left == nil {
			return Aux(root.Right, path)
		}
		if root.Right == nil {
			return Aux(root.Left, path)
		}
		return Smaller(Aux(root.Left, path), Aux(root.Right, path))
	}
	return Reverse(path)
}

func smallestFromLeaf(root *TreeNode) (ret string) {
	if root != nil {
		ret = Aux(root, "")
	}
	return ret
}

func main() {}
