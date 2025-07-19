// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/
package main

import "strings"

type (
	Trie struct{ head TrieNode }

	TrieNode struct {
		children map[string]*TrieNode
		end      bool
	}
)

func trieNode() *TrieNode {
	return &TrieNode{children: make(map[string]*TrieNode)}
}

func (node *TrieNode) get() (paths []string) {
	if node.end {
		return []string{""}
	}
	for key, newNode := range node.children {
		for _, sub := range newNode.get() {
			paths = append(paths, "/"+key+sub)
		}
	}
	return paths
}

func (trie *Trie) insert(sentence []string) {
	node := &trie.head
	for _, word := range sentence {
		if _, found := node.children[word]; !found {
			node.children[word] = trieNode()
		}
		node = node.children[word]
	}
	node.end = true
}

func removeSubfolders(folder []string) []string {
	trie := Trie{head: *trieNode()}
	for _, path := range folder {
		trie.insert(strings.Split(path, "/")[1:])
	}
	return trie.head.get()
}

func main() {}
