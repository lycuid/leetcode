// https://leetcode.com/problems/hand-of-straights/
package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if groupSize == 1 {
		return true
	}

	sort.Ints(hand)
	type Tuple struct{ value, count int }
	var freq []Tuple

	for _, value := range hand {
		if n := len(freq); n > 0 && freq[n-1].value == value {
			freq[n-1].count++
		} else {
			freq = append(freq, Tuple{value, 1})
		}
	}
	for len(freq) >= groupSize {
		tup := freq[0]
		for i := 1; i < groupSize; i++ {
			if freq[i].value != freq[i-1].value+1 || freq[i].count < tup.count {
				return false
			}
			freq[i].count -= tup.count
		}
		for freq = freq[1:]; len(freq) > 0 && freq[0].count == 0; {
			freq = freq[1:]
		}
	}
	return len(freq) == 0
}

func main() {}
