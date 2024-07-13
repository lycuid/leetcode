// https://leetcode.com/problems/robot-collisions/
package main

import "sort"

type Robot struct{ pos, hp, dir, index int }

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	var (
		stack  []Robot
		robots = make([]Robot, len(positions))
	)
	for i := range robots {
		robots[i] = Robot{positions[i], healths[i], int(directions[i]), i}
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].pos < robots[j].pos
	})
	for _, robot := range robots {
		for n := len(stack); n > 0 && stack[n-1].dir > robot.dir && stack[n-1].hp < robot.hp; n-- {
			stack, robot.hp = stack[:n-1], robot.hp-1
		}
		if n := len(stack); n > 0 {
			if stack[n-1].dir > robot.dir && stack[n-1].hp >= robot.hp {
				if stack[n-1].hp == robot.hp {
					stack = stack[:n-1]
				} else {
					stack[n-1].hp--
				}
				continue
			}
		}
		stack = append(stack, robot)
	}
	sort.Slice(stack, func(i, j int) bool {
		return stack[i].index < stack[j].index
	})
	hps := make([]int, len(stack))
	for i := range stack {
		hps[i] = stack[i].hp
	}
	return hps
}

func main() {}
