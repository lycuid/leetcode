// https://leetcode.com/problems/find-duplicate-subtrees/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Depth = int
type Value = int

func Matches(fst, snd *TreeNode) bool {
	return fst == snd || (fst != nil && snd != nil && fst.Val == snd.Val && Matches(fst.Left, snd.Left) && Matches(fst.Right, snd.Right))
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Aux(root *TreeNode, cache map[Depth]map[Value][][]*TreeNode) (depth int) {
	if root != nil {
		if depth = Max(Aux(root.Left, cache), Aux(root.Right, cache)) + 1; cache[depth] == nil {
			cache[depth] = make(map[Value][][]*TreeNode)
		}
		for i, nodes := range cache[depth][root.Val] {
			if len(nodes) > 0 && Matches(nodes[0], root) {
				cache[depth][root.Val][i] = append(cache[depth][root.Val][i], root)
				goto EXIT
			}
		}
		cache[depth][root.Val] = append(cache[depth][root.Val], []*TreeNode{root})
	}
EXIT:
	return depth
}

func findDuplicateSubtrees(root *TreeNode) (ret []*TreeNode) {
	cache := make(map[Depth]map[Value][][]*TreeNode)
	Aux(root, cache)
	for depth := range cache {
		for value := range cache[depth] {
			for _, nodes := range cache[depth][value] {
				if len(nodes) > 1 {
					ret = append(ret, nodes[0])
				}
			}
		}
	}
	return ret
}

func main() {}
