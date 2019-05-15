// https://leetcode.com/problems/roman-to-integer/

package main

func romanToInt(s string) int {
	romans := map[string]int{
		"":  0,
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	preRomans := map[int]int{
		5:    1,
		10:   1,
		50:   10,
		100:  10,
		500:  100,
		1000: 100,
	}

	if len(s) <= 1 {
		return romans[s]
	}

	result := 0
	add := true

	for i := len(s) - 1; i > 0; i-- {
		number := romans[string(s[i])]

		if add {
			result += number
		} else {
			result -= number
		}

		if next, ok := preRomans[number]; ok {
			if next == romans[string(s[i-1])] {
				add = false
				continue
			}
		}

		add = true
	}

	if add {
		result += romans[string(s[0])]
	} else {
		result -= romans[string(s[0])]
	}

	return result
}

func main() {}
