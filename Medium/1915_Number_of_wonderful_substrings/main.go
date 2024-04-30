// https://leetcode.com/problems/number-of-wonderful-substrings/
package main

func wonderfulSubstrings(word string) (ret int64) {
	for i, prefix, cache := 0, 0, [1 << 10]int64{1}; i < len(word); i++ {
		ch := int64(word[i] - 'a')
		prefix ^= 1 << ch
		ret += cache[prefix]
		for j := 0; j < 10; j++ {
			ret += cache[prefix^(1<<j)]
		}
		cache[prefix]++
	}
	return ret
}

func main() {}
