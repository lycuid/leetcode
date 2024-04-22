// https://leetcode.com/problems/open-the-lock/
package main

type Tuple [4]int

func FromString(str string) Tuple {
	return Tuple{
		int(str[0] - '0'),
		int(str[1] - '0'),
		int(str[2] - '0'),
		int(str[3] - '0'),
	}
}

type Trie struct{ inner [10]*Trie }

func (trie *Trie) Insert(tup Tuple) {
	node := trie
	for _, num := range tup {
		if node.inner[num] == nil {
			node.inner[num] = &Trie{}
		}
		node = node.inner[num]
	}
}

func (trie *Trie) Has(tup Tuple) bool {
	node := trie
	for _, num := range tup {
		if node.inner[num] == nil {
			return false
		}
		node = node.inner[num]
	}
	return true
}

func openLock(deadends []string, target string) int {
	done, path, src, dst := Trie{}, -1, FromString("0000"), FromString(target)
	for _, deadend := range deadends {
		done.Insert(FromString(deadend))
	}
	stack, parent, children := [2][]Tuple{{}, {src}}, 0, 1
	for len(stack[parent]) > 0 || len(stack[children]) > 0 {
		parent, children, path = children, parent, path+1
		for ; len(stack[parent]) > 0; stack[parent] = stack[parent][1:] {
			if node := stack[parent][0]; !done.Has(node) {
				if done.Insert(node); node == dst {
					return path
				}
				for i := 0; i < 4; i++ {
					for _, j := range [2]int{-1, 1} {
						child := node
						child[i] = (child[i] + j + 10) % 10
						stack[children] = append(stack[children], child)
					}
				}
			}
		}
	}
	return -1
}

func main() {}
