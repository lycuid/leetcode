// https://leetcode.com/problems/design-add-and-search-words-data-structure/
package main

type WordDictionary struct {
	inner [26]*WordDictionary
	end   bool
}

func Constructor() WordDictionary { return WordDictionary{end: true} }

func (this *WordDictionary) AddWord(word string) {
	for _, ch := range word {
		if this.inner[ch-'a'] == nil {
			this.inner[ch-'a'] = new(WordDictionary)
		}
		this = this.inner[ch-'a']
	}
	this.end = true
}

func (this *WordDictionary) Search(word string) (ret bool) {
	for i, ch := range word {
		if ch == '.' {
			for _, node := range this.inner {
				if node != nil && node.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		if this.inner[ch-'a'] == nil {
			return false
		}
		this = this.inner[ch-'a']
	}
	return this.end
}

func main() {}
