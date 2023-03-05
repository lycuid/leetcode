// https://leetcode.com/problems/jump-game-iv/
package main

type Tuple struct{ index, value, weight int }

func minJumps(nums []int) int {
	index_of, n, VISITED := make(map[int][]int), len(nums), -int(1e9+7)
	for i, num := range nums {
		index_of[num] = append(index_of[num], i)
	}
	for q := []Tuple{{0, nums[0], 0}}; len(q) > 0; q = q[1:] {
		p := q[0]
		if p.index == n-1 {
			return p.weight
		}
		if i := p.index - 1; i >= 0 && nums[i] != VISITED {
			q, nums[i] = append(q, Tuple{i, nums[i], p.weight + 1}), VISITED
		}
		if i := p.index + 1; i < n && nums[i] != VISITED {
			q, nums[i] = append(q, Tuple{i, nums[i], p.weight + 1}), VISITED
		}
		if indices, found := index_of[p.value]; found {
			for _, i := range indices {
				if nums[i] != VISITED {
					q, nums[i] = append(q, Tuple{i, nums[i], p.weight + 1}), VISITED
				}
			}
			delete(index_of, p.value)
		}
	}
	return 0
}

func main() {}
