// https://leetcode.com/problems/insert-delete-getrandom-o1/
package main

import "math/rand"

type RandomizedSet struct {
	cache map[int]int
	list  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{cache: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) (found bool) {
	if _, found = this.cache[val]; !found {
		this.list = append(this.list, val)
		this.cache[val] = len(this.list) - 1
	}
	return !found
}

func (this *RandomizedSet) Remove(val int) (found bool) {
	if _, found = this.cache[val]; found {
		n, index := len(this.list)-1, this.cache[val]
		this.list[index], this.list = this.list[n], this.list[:n]
		if n != index {
			this.cache[this.list[index]] = index
		}
		delete(this.cache, val)
	}
	return found
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

func main() {}
