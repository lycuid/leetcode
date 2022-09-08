// https://leetcode.com/problems/word-ladder-ii/
package main

type PriorityQueue struct {
	inner []string
}

func MakePQ() PriorityQueue {
	return PriorityQueue{make([]string, 1)}
}

func (this *PriorityQueue) Push(item string, weight map[string]int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && weight[this.inner[i]] < weight[this.inner[i/2]]; i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop(weight map[string]int) string {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && weight[this.inner[j+1]] < weight[this.inner[j]] {
			j++
		}
		if weight[this.inner[i]] < weight[this.inner[j]] {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func (this PriorityQueue) Empty() bool {
	return len(this.inner) == 1
}

func isChild(w1, w2 string) bool {
	for i := range w1 {
		if w1[i] != w2[i] {
			return w1[:i] == w2[:i] && w1[i+1:] == w2[i+1:]
		}
	}
	return false
}

func Path(parent string, graph map[string][]string, weight map[string]int) (ret [][]string) {
	if weight[parent] == 0 {
		return [][]string{{parent}}
	}
	for _, child := range graph[parent] {
		if weight[child] == weight[parent]-1 {
			for _, path := range Path(child, graph, weight) {
				ret = append(ret, append([]string{parent}, path...))
			}
		}
	}
	return ret
}

func Entries(s string, words []string, weight map[string]int) (ret []string) {
	var min_weight int = 1e5
	for _, word := range words {
		if isChild(s, word) && weight[word] < min_weight {
			min_weight = weight[word]
		}
	}
	for _, word := range words {
		if isChild(s, word) && weight[word] == min_weight {
			ret = append(ret, word)
		}
	}
	return ret
}

func findLadders(src string, dest string, words []string) (ret [][]string) {
	graph := make(map[string][]string)
	weight, visited := make(map[string]int), make(map[string]bool)
	for _, word1 := range words {
		weight[word1] = 1e5
		for _, word2 := range words {
			if isChild(word1, word2) {
				graph[word1] = append(graph[word1], word2)
			}
		}
	}
	pq := MakePQ()
	pq.Push(dest, weight)
	weight[dest], visited[dest] = 0, true
	for !pq.Empty() {
		parent := pq.Pop(weight)
		for _, child := range graph[parent] {
			if !visited[child] {
				pq.Push(child, weight)
				if visited[child] = true; weight[child] > weight[parent]+1 {
					weight[child] = weight[parent] + 1
				}
			}
		}
	}
	for _, entry := range Entries(src, words, weight) {
		for _, path := range Path(entry, graph, weight) {
			ret = append(ret, append([]string{src}, path...))
		}
	}
	return ret
}

func main() {}
