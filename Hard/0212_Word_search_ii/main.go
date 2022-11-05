// https://leetcode.com/problems/word-search-ii/
package main

type Trie struct {
	end      bool
	value    string
	children [26]*Trie
}

func (this *Trie) Insert(value string) {
	root := this
	for _, ch := range value {
		if root.children[ch-'a'] == nil {
			root.children[ch-'a'] = &Trie{}
		}
		root = root.children[ch-'a']
	}
	root.end, root.value = true, value
}

func Dfs(trie *Trie, y, x int, board [][]byte, visited [][]bool, ret *[]string) {
	if y < 0 || y >= len(board) || x < 0 || x >= len(board[0]) || visited[y][x] {
		return
	}
	node := trie.children[board[y][x]-'a']
	if node == nil {
		return
	}
	if node.end {
		*ret, node.end = append(*ret, node.value), false
	}
	visited[y][x] = true
	for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		Dfs(node, y+d[0], x+d[1], board, visited, ret)
	}
	visited[y][x] = false
}

func findWords(board [][]byte, words []string) (ret []string) {
	trie, visited := Trie{}, make([][]bool, len(board))
	for _, word := range words {
		trie.Insert(word)
	}
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for y := range board {
		for x := range board[0] {
			Dfs(&trie, y, x, board, visited, &ret)
		}
	}
	return ret
}

func main() {}
