// https://leetcode.com/problems/walking-robot-simulation-ii/
package main

const (
	East = iota
	North
	West
	South
)

type Robot struct {
	w, h, dir int
	pos       []int
}

func Constructor(width int, height int) Robot {
	return Robot{w: width, h: height, pos: make([]int, 2)}
}

func (this *Robot) Step(num int) {
	num = num % (2*(this.w+this.h) - 4)
	if num == 0 {
		if this.pos[0] == 0 && this.pos[1] == 0 {
			this.dir = South
		}
	}
	for num > 0 {
		switch this.dir {
		case East:
			if this.pos[0]+num < this.w {
				this.pos[0], num = this.pos[0]+num, 0
			} else {
				this.pos[0], num = this.w-1, num-(this.w-1-this.pos[0])
				this.dir = (this.dir + 1) % 4
			}
		case North:
			if this.pos[1]+num < this.h {
				this.pos[1], num = this.pos[1]+num, 0
			} else {
				this.pos[1], num = this.h-1, num-(this.h-1-this.pos[1])
				this.dir = (this.dir + 1) % 4
			}
		case West:
			if this.pos[0]-num >= 0 {
				this.pos[0], num = this.pos[0]-num, 0
			} else {
				this.pos[0], num = 0, num-this.pos[0]
				this.dir = (this.dir + 1) % 4
			}
		case South:
			if this.pos[1]-num >= 0 {
				this.pos[1], num = this.pos[1]-num, 0
			} else {
				this.pos[1], num = 0, num-this.pos[1]
				this.dir = (this.dir + 1) % 4
			}
		}
	}
}

func (this *Robot) GetPos() []int {
	return this.pos
}

func (this *Robot) GetDir() string {
	switch this.dir {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return ""
}

func main() {}
