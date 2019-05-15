// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

func lengthOfLongestSubstring(s string) int {
	length := 0
	result := 0
	visited := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		b := s[i]
		if index, ok := visited[b]; ok {
			if length > result {
				result = length
			}
			length = 0
			visited = map[byte]int{}
			i = index

			continue
		}

		visited[b] = i
		length++
	}

	if length > result {
		result = length
	}
	return result
}

func main() {}
