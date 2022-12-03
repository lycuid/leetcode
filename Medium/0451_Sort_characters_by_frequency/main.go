// https://leetcode.com/problems/sort-characters-by-frequency/
package main

func frequencySort(s string) (ret string) {
	var freq [128]int
	var max_freq int
	for _, ch := range s {
		if freq[ch]++; freq[ch] >= max_freq {
			max_freq = freq[ch] + 1
		}
	}
	bytes_count := make([][]byte, max_freq+1)
	for ch, f := range freq {
		if f > 0 {
			bytes_count[f] = append(bytes_count[f], byte(ch))
		}
	}
	for i := max_freq; i >= 0; i-- {
		for _, ch := range bytes_count[i] {
			for j := 0; j < i; j++ {
				ret += string(ch)
			}
		}
	}
	return ret
}

func main() {}
