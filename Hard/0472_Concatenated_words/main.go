// https://leetcode.com/problems/concatenated-words/
package main

func IsPrefix(small, big string) bool {
	for l, r := 0, len(small)-1; l <= r; l, r = l+1, r-1 {
		if small[l] != big[l] || small[r] != big[r] {
			return false
		}
	}
	return true
}

func Aux(big string, cache [][]string) bool {
	for i := len(big); i > 0; i-- {
		for _, small := range cache[i] {
			if IsPrefix(small, big) && Aux(big[i:], cache) {
				return true
			}
		}
	}
	return len(big) == 0
}

func findAllConcatenatedWordsInADict(words []string) (ret []string) {
	var max int
	for _, word := range words {
		if max < len(word) {
			max = len(word)
		}
	}
	cache := make([][]string, max+1)
	for _, word := range words {
		cache[len(word)] = append(cache[len(word)], word)
	}
	for l := max; l >= 0; l-- {
		for _, big := range cache[l] {
			for i := len(big) - 1; i > 0; i-- {
				for _, small := range cache[i] {
					if IsPrefix(small, big) && Aux(big[i:], cache) {
						ret, i = append(ret, big), 0
						break
					}
				}
			}
		}
	}
	return ret
}

func main() {}
