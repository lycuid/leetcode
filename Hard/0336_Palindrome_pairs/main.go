// https://leetcode.com/problems/palindrome-pairs/
package main

func palindromePairs(words []string) (ret [][]int) {
	indexOf := make(map[string]int)
	for w, word := range words {
		tmp := make([]byte, len(word))
		for i := range word {
			tmp[len(tmp)-1-i] = word[i]
		}
		indexOf[string(tmp)] = w
	}
	for w, word := range words {
		l := len(word)
		for n := 0; n < l; n++ {
			for _, s := range [][]int{{n, n}, {n, n + 1}} {
				i, j := s[0], s[1]
				for ; i >= 0 && j < l && word[i] == word[j]; i, j = i-1, j+1 {
				}
				if i < 0 {
					if index, found := indexOf[word[j:]]; found && index != w {
						ret = append(ret, []int{index, w})
					}
				}
				if j >= l {
					if index, found := indexOf[word[:i+1]]; found && index != w {
						ret = append(ret, []int{w, index})
					}
				}
			}
		}
	}
	return ret
}

func main() {}
