// https://leetcode.com/problems/insert-delete-getrandom-o1/
package main

import "math/rand"

type RandomizedSet struct {
	inner   []int
	indexOf map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 1), make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.indexOf[val] > 0 {
		return false
	}
	this.inner, this.indexOf[val] = append(this.inner, val), len(this.inner)
	return true
}

func (this *RandomizedSet) Remove(val1 int) bool {
	if this.indexOf[val1] == 0 {
		return false
	}
	i, j := this.indexOf[val1], len(this.inner)-1
	val2 := this.inner[j]
	this.inner[i], this.inner = this.inner[j], this.inner[:j]
	if this.indexOf[val1], this.indexOf[val2] = 0, i; val1 == val2 {
		this.indexOf[val2] = 0
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.inner[rand.Intn(len(this.inner)-1)+1]
}

func main() {}
