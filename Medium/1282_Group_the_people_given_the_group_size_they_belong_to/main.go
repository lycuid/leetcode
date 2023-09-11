// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
package main

func groupThePeople(groupSizes []int) (ret [][]int) {
	cache := make([][]int, len(groupSizes)+1)
	for i, num := range groupSizes {
		if cache[num] = append(cache[num], i); len(cache[num]) == num {
			ret, cache[num] = append(ret, cache[num]), []int{}
		}
	}
	return ret
}

func main() {}
