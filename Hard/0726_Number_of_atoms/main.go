// https://leetcode.com/problems/number-of-atoms/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func parse_atom(stream string) (string, string) {
	i := 1
	for i < len(stream) && stream[i] >= 'a' && stream[i] <= 'z' {
		i++
	}
	return string(stream[:i]), stream[i:]
}

func parse_num(stream string) (int, string) {
	if len(stream) > 0 {
		var i int
		for i < len(stream) && stream[i] >= '0' && stream[i] <= '9' {
			i++
		}
		if num, err := strconv.Atoi(string(stream[:i])); err == nil {
			return num, stream[i:]
		}
	}
	return 1, stream
}

func parse(stream string, mul int) map[string]int {
	atoms := make(map[string]int)
	for len(stream) > 0 {
		var count int
		if ch := stream[0]; ch >= 'A' && ch <= 'Z' {
			name, rest := parse_atom(stream)
			count, stream = parse_num(rest)
			atoms[name] += count * mul
		} else if ch == '(' {
			i := 1
			for depth := 1; i < len(stream); i++ {
				switch stream[i] {
				case '(':
					depth++
				case ')':
					depth--
				}
				if depth == 0 {
					break
				}
			}
			sub := stream[1:i]
			count, stream = parse_num(stream[i+1:])
			for name, count := range parse(sub, count) {
				atoms[name] += count * mul
			}
		}
	}
	return atoms
}

func countOfAtoms(formula string) (ret string) {
	type Atom struct {
		name  string
		count int
	}
	cache := parse(formula, 1)
	atoms := make([]Atom, len(cache))
	for name, count := range cache {
		atoms = append(atoms, Atom{name, count})
	}
	sort.Slice(atoms, func(i, j int) bool {
		return atoms[i].name < atoms[j].name
	})
	for _, atom := range atoms {
		if ret += atom.name; atom.count > 1 {
			ret += fmt.Sprintf("%d", atom.count)
		}
	}
	return ret
}

func main() {}
