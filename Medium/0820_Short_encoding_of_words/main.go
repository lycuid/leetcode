// https://leetcode.com/problems/short-encoding-of-words/
package main

type TrieNode struct {
	Children [26]*TrieNode
}

func (this *TrieNode) Dfs() (ret []int) {
	for _, child := range this.Children {
		if child != nil {
			ns := child.Dfs()
			for _, n := range ns {
				ret = append(ret, 1+n)
			}
		}
	}
	if len(ret) == 0 {
		ret = []int{1}
	}
	return ret
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
}

func (this *Trie) Dfs() (ret int) {
	ns := this.root.Dfs()
	for _, n := range ns {
		ret += n
	}
	return ret
}

func Reverse(word string) string {
	n, chars := len(word), []byte(word)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}
	return string(chars)
}

func minimumLengthEncoding(words []string) int {
	trie := Trie{new(TrieNode)}
	for _, word := range words {
		trie.Insert(Reverse(word))
	}
	return trie.Dfs()
}

func main() {}
