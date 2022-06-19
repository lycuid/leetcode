// https://leetcode.com/problems/search-suggestions-system/
package main

import "fmt"

type TrieNode struct {
	End      bool
	Children [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, ch := range word {
		at := ch - 'a'
		if current.Children[at] == nil {
			current.Children[at] = new(TrieNode)
		}
		current = current.Children[at]
	}
	current.End = true
}

func Take3(node *TrieNode) (ret []string) {
	if node.End {
		ret = append(ret, "")
	}
	for i := 0; i < 26 && len(ret) < 3; i++ {
		ch, next := rune(i+'a'), node.Children[i]
		if next != nil {
			words := Take3(next)
			if len(words) == 0 && !node.End {
				ret = append(ret, fmt.Sprintf("%c", ch))
			}
			for i := 0; i < len(words); i++ {
				ret = append(ret, fmt.Sprintf("%c%s", ch, words[i]))
			}
		}
	}
	if len(ret) > 3 {
		return ret[:3]
	}
	return ret
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Trie{root: new(TrieNode)}
	ret := make([][]string, len(searchWord))
	for _, word := range products {
		trie.Insert(word)
	}
	current := trie.root
	for i, ch := range searchWord {
		at := ch - 'a'
		if current.Children[at] == nil {
			break
		}
		ret[i] = Take3(current.Children[at])
		for w, word := range ret[i] {
			ret[i][w] = fmt.Sprintf("%s%s", searchWord[:i+1], word)
		}
		current = current.Children[at]
	}
	return ret
}

func main() {}
