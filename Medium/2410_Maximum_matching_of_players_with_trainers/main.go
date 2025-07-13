// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/
package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) (count int) {
	sort.Slice(players, func(i, j int) bool {
		return players[i] > players[j]
	})
	sort.Slice(trainers, func(i, j int) bool {
		return trainers[i] > trainers[j]
	})

	for ; len(players) > 0 && len(trainers) > 0; players = players[1:] {
		if players[0] <= trainers[0] {
			trainers, count = trainers[1:], count+1
		}
	}

	return count
}

func main() {}
