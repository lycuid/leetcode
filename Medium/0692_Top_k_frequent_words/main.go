// https://leetcode.com/problems/top-k-frequent-words/
package main

import "strings"

type PriorityQueue struct {
	inner []string
}

func (this *PriorityQueue) Init()      { this.inner = make([]string, 1) }
func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }
func (this PriorityQueue) Less(i, j int) bool {
	return strings.Compare(this.inner[i], this.inner[j]) == -1
}

func (this *PriorityQueue) Push(item string) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.Less(i, i/2); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() string {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.Less(j+1, j) {
			j++
		}
		if this.Less(i, j) {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func topKFrequent(words []string, k int) (ret []string) {
	count, freq := make(map[string]int), make([]PriorityQueue, len(words)+1)
	for i := range freq {
		freq[i].Init()
	}
	for _, word := range words {
		count[word]++
	}
	for _, word := range words {
		if count[word] > 0 {
			freq[count[word]].Push(word)
			count[word] = 0
		}
	}
	for i := len(freq) - 1; i >= 0 && k > 0; i-- {
		for ; !freq[i].Empty() && k > 0; k-- {
			ret = append(ret, freq[i].Pop())
		}
	}
	return ret
}

func main() {}
