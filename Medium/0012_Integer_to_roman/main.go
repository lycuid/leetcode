// https://leetcode.com/problems/integer-to-roman/

package main

import "bytes"

func getUnitNumber(n int) int {
	num := 1
	for n%10 == 0 {
		n /= 10
		num *= 10
	}
	return num
}

func toRoman(num int) string {
	romans := map[int]string{
		0:    "",
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	if n, ok := romans[num]; ok {
		return n
	}

	unitNumber := getUnitNumber(num)
	unitBounds := []int{unitNumber * 5, unitNumber * 10}

	var roman bytes.Buffer

	if num == unitBounds[0]-unitNumber {
		return romans[unitNumber] + romans[unitBounds[0]]
	}

	if num < unitBounds[0] {
		for i := 0; i < num/unitNumber; i++ {
			roman.WriteString(romans[unitNumber])
		}
		return roman.String()
	}

	if num == unitBounds[1]-unitNumber {
		return romans[unitNumber] + romans[unitBounds[1]]
	}

	if num < unitBounds[1] {
		roman.WriteString(romans[unitBounds[0]])
		for i := 0; i < (num-unitBounds[0])/unitNumber; i++ {
			roman.WriteString(romans[unitNumber])
		}
		return roman.String()
	}

	return ""
}

func intToRoman(num int) (roman string) {
	unit := 10
	for num > 0 {
		n := num % unit
		roman = toRoman(n) + roman
		num -= n
		unit *= 10
	}
	return
}

func main() {}
