// https://leetcode.com/problems/string-to-integer-atoi/

package main

// Predicate fn
type Predicate func(rune) bool

func toNumber(ns []int) (number int) {
	for _, n := range ns {
		number = number*10 + n
	}
	return
}

func takeWhile(fn Predicate, s string) string {
	for i, r := range s {
		if !fn(r) {
			return s[:i]
		}
	}
	return s
}

func dropWhile(fn Predicate, s string) string {
	for i, r := range s {
		if !fn(r) {
			return s[i:]
		}
	}
	return ""
}

func toSigned(s string) (int, string) {
	if len(s) < 1 {
		return 1, ""
	}

	signs := map[rune]int{43: 1, 45: -1}
	head := rune(s[0])
	if sign, ok := signs[head]; ok {
		return sign, s[1:]
	}

	return 1, s
}

func mapNumber(str string) []int {
	numbers := make([]int, len(str))
	for i, r := range str {
		numbers[i] = int(r) - 48
	}
	return numbers
}

func checkAndConvert(ns, inf []int) int {
	lns := len(ns)
	linf := len(inf)

	if lns > linf {
		return toNumber(inf)
	}

	if lns < linf {
		return toNumber(ns)
	}

	ns1 := toNumber(ns[:lns])
	inf1 := toNumber(inf[:linf])
	if ns1 > inf1 || (ns1 == inf1 && toNumber([]int{ns[lns-1]}) >= toNumber([]int{inf[linf-1]})) {
		return toNumber(inf)
	}

	return toNumber(ns)
}

func myAtoi(str string) int {
	maxIntArr := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	minIntArr := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 8}

	// trim beginning.
	str = dropWhile(func(r rune) bool { return r == 32 }, str)

	sign, str := toSigned(str)

	// trim for zeroes
	str = dropWhile(func(r rune) bool { return r == 48 }, str)
	// trim till invalid.
	str = takeWhile(func(r rune) bool { return r >= 48 && r <= 57 }, str)

	numberList := mapNumber(str)
	if sign < 0 {
		return sign * checkAndConvert(numberList, minIntArr)
	}

	return sign * checkAndConvert(numberList, maxIntArr)
}

func main() {}
