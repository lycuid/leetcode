// https://leetcode.com/problems/extra-characters-in-a-string/
package main

type TrieNode struct {
	children [26]*TrieNode
	end      bool
}

type Trie struct{ head TrieNode }

func (trie *Trie) Insert(str string) {
	node := &trie.head
	for i := range str {
		ch := str[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = new(TrieNode)
		}
		node = node.children[ch]
	}
	node.end = true
}

func minExtraChar(s string, dictionary []string) int {
	var (
		n     = len(s)
		cache = make([]int, n+1)
		trie  Trie
	)

	for i := range cache {
		cache[i] = i
	}
	for _, word := range dictionary {
		trie.Insert(word)
	}

	for i := 0; i < n; i++ {
		cache[i+1] = min(cache[i+1], cache[i]+1)
		node := &trie.head
		for j := i; j < n; j++ {
			ch := s[j] - 'a'
			if node.children[ch] == nil {
				break
			}
			if node = node.children[ch]; node.end {
				cache[j+1] = min(cache[j+1], cache[i])
			}
		}
	}
	return cache[n]
}

func main() {}
