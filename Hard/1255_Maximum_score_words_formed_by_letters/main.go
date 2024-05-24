// https://leetcode.com/problems/maximum-score-words-formed-by-letters/
package main

type Freq [26]int

func (f0 Freq) FitsIn(f1 Freq) bool {
	for i, val := range f1 {
		if f0[i] > val {
			return false
		}
	}
	return true
}

func (f0 *Freq) Subtract(f1 Freq) *Freq {
	for i, val := range f1 {
		f0[i] -= val
	}
	return f0
}

func (f0 *Freq) Add(f1 Freq) *Freq {
	for i, val := range f1 {
		f0[i] += val
	}
	return f0
}

func (f Freq) CalculateScore(points []int) (score int) {
	for ch, cnt := range f {
		score += points[ch] * cnt
	}
	return score
}

func FromString(word string) (freq Freq) {
	for _, ch := range word {
		freq[ch-'a']++
	}
	return freq
}

func Aux(words []Freq, stash Freq, points []int) (ret int) {
	for i, word := range words {
		if word.FitsIn(stash) {
			score := Aux(words[:i], *stash.Subtract(word), points)
			if total := score + word.CalculateScore(points); total > ret {
				ret = total
			}
			stash.Add(word)
		}
	}
	return ret
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	var (
		stash = FromString(string(letters))
		freq  = make([]Freq, len(words))
	)
	for i, word := range words {
		freq[i] = FromString(word)
	}
	return Aux(freq, stash, score)
}

func main() {}
