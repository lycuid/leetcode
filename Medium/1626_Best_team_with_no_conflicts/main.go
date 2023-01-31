// https://leetcode.com/problems/best-team-with-no-conflicts/
package main

import "sort"

type Tuple struct {
	score, age int
	done       bool
}

func Aux(i int, cache []Tuple) int {
	if !cache[i].done {
		cache[i].done = true
		for j, init := i+1, cache[i].score; j < len(cache); j++ {
			if cache[i].age == cache[j].age || init <= cache[j].score {
				if score := init + Aux(j, cache); score > cache[i].score {
					cache[i].score = score
				}
			}
		}
	}
	return cache[i].score
}

func bestTeamScore(scores []int, ages []int) (ret int) {
	cache := make([]Tuple, len(scores))
	for i := range cache {
		cache[i] = Tuple{score: scores[i], age: ages[i]}
	}
	sort.Slice(cache, func(i, j int) bool {
		if cache[i].age == cache[j].age {
			return cache[i].score < cache[j].score
		}
		return cache[i].age < cache[j].age
	})
	for i := range cache {
		if score := Aux(i, cache); score > ret {
			ret = score
		}
	}
	return ret
}

func main() {}
