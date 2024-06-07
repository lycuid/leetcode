// https://leetcode.com/problems/replace-words/
package main

import "strings"

type TrieNode struct {
	children [26]*TrieNode
	end      bool
}

type Trie struct{ head TrieNode }

func (trie *Trie) Insert(word string) {
	node := &trie.head
	for i := range word {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{}
		}
		node = node.children[ch]
	}
	node.end = true
}

func (trie Trie) GetPrefix(word string) string {
	node := &trie.head
	for i := range word {
		ch := word[i] - 'a'
		if node.end {
			return word[:i]
		}
		if node.children[ch] == nil {
			break
		}
		node = node.children[ch]
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	var trie Trie
	words := strings.Split(sentence, " ")
	for _, word := range dictionary {
		trie.Insert(word)
	}
	for i, word := range words {
		words[i] = trie.GetPrefix(word)
	}
	return strings.Join(words, " ")
}

func main() {}
