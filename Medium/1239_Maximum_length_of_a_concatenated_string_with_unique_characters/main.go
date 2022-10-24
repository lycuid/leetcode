// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
package main

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Aux(mask, i int, words []string, masks []int) (ret int) {
	for j := i; j < len(words); j++ {
		if mask&masks[j] == 0 {
			ret = Max(ret, len(words[j])+Aux(mask|masks[j], j+1, words, masks))
		}
	}
	return ret
}

func maxLength(arr []string) (ret int) {
	var (
		words []string
		masks []int
	)
MAINLOOP:
	for _, word := range arr {
		mask := 0
		for _, ch := range word {
			if (mask>>(ch-'a'))&1 != 0 {
				continue MAINLOOP
			}
			mask |= 1 << (ch - 'a')
		}
		words, masks = append(words, word), append(masks, mask)
	}
	for i := range words {
		ret = Max(ret, len(words[i])+Aux(masks[i], i+1, words, masks))
	}
	return ret
}

func main() {}
