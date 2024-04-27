// https://leetcode.com/problems/freedom-trail/
package main

import "math"

type Tuple struct{ index, dist int }

type Array struct {
	items       []Tuple
	Length, Min int
}

func MakeArray(capacity int) Array {
	return Array{items: make([]Tuple, capacity), Min: math.MaxInt}
}

func (arr *Array) Clear() { arr.Length, arr.Min = 0, math.MaxInt }

func (arr *Array) Append(item Tuple) {
	arr.items[arr.Length], arr.Length = item, arr.Length+1
	if item.dist < arr.Min {
		arr.Min = item.dist
	}
}

func findRotateSteps(ring string, key string) int {
	var indices [26][]int
	for i := range ring {
		indices[ring[i]-'a'] = append(indices[ring[i]-'a'], i)
	}

	MinDist := func(i, j int) int {
		if j < i {
			i, j = j, i
		}
		if n := i + len(ring) - j; n < j-i {
			return n
		}
		return j - i
	}

	arr := []Array{MakeArray(len(ring)), MakeArray(len(ring))}
	prev, curr := 0, 1

	arr[curr].Append(Tuple{0, 0})
	for i := range key {
		prev, curr = curr, prev
		arr[curr].Clear()
		for _, index := range indices[key[i]-'a'] {
			tup := Tuple{dist: math.MaxInt}
			for j := 0; j < arr[prev].Length; j++ {
				t1 := arr[prev].items[j]
				if dist := t1.dist + MinDist(t1.index, index) + 1; dist < tup.dist {
					tup.dist, tup.index = dist, index
				}
			}
			arr[curr].Append(tup)
		}
	}
	return arr[curr].Min
}

func main() {}
