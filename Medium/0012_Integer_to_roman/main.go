// https://leetcode.com/problems/integer-to-roman/
package main

func intToRoman(num int) string {
	switch num {
	case 1:
		return "I"
	case 4:
		return intToRoman(1) + intToRoman(5)
	case 5:
		return "V"
	case 9:
		return intToRoman(1) + intToRoman(10)
	case 10:
		return "X"
	case 40:
		return intToRoman(10) + intToRoman(50)
	case 50:
		return "L"
	case 90:
		return intToRoman(10) + intToRoman(100)
	case 100:
		return "C"
	case 400:
		return intToRoman(100) + intToRoman(500)
	case 500:
		return "D"
	case 900:
		return intToRoman(100) + intToRoman(1000)
	case 1000:
		return "M"
	}
	if num > 1 && num < 4 {
		return intToRoman(1) + intToRoman(num-1)
	} else if num > 4 && num < 5 {
		return intToRoman(4) + intToRoman(num-4)
	} else if num > 5 && num < 9 {
		return intToRoman(5) + intToRoman(num-5)
	} else if num > 9 && num < 10 {
		return intToRoman(9) + intToRoman(num-9)
	} else if num > 10 && num < 40 {
		return intToRoman(10) + intToRoman(num-10)
	} else if num > 40 && num < 50 {
		return intToRoman(40) + intToRoman(num-40)
	} else if num > 50 && num < 90 {
		return intToRoman(50) + intToRoman(num-50)
	} else if num > 90 && num < 100 {
		return intToRoman(90) + intToRoman(num-90)
	} else if num > 100 && num < 400 {
		return intToRoman(100) + intToRoman(num-100)
	} else if num > 400 && num < 500 {
		return intToRoman(400) + intToRoman(num-400)
	} else if num > 500 && num < 900 {
		return intToRoman(500) + intToRoman(num-500)
	} else if num > 900 && num < 1000 {
		return intToRoman(900) + intToRoman(num-900)
	} else if num > 1000 {
		return intToRoman(1000) + intToRoman(num-1000)
	}
	return ""
}

func main() {}
