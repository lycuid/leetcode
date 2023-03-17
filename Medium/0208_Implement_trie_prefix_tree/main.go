// https://leetcode.com/problems/implement-trie-prefix-tree/
package main

type Trie struct {
	inner [26]*Trie
	end   bool
}

func Constructor() Trie { return Trie{end: true} }

func (this *Trie) Last(word string) *Trie {
	for _, ch := range word {
		if this.inner[ch-'a'] == nil {
			return nil
		}
		this = this.inner[ch-'a']
	}
	return this
}

func (this *Trie) Insert(word string) {
	for _, ch := range word {
		if this.inner[ch-'a'] == nil {
			this.inner[ch-'a'] = &Trie{}
		}
		this = this.inner[ch-'a']
	}
	this.end = true
}

func (this *Trie) Search(word string) bool {
	last := this.Last(word)
	return last != nil && last.end
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.Last(prefix) != nil
}

func main() {}
