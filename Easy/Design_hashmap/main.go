// https://leetcode.com/problems/design-hashmap/
package main

type MyHashMap struct {
	hashmap []int
}

func Constructor() MyHashMap {
	hashmap := make([]int, 1e6+1)
	for i := range hashmap {
		hashmap[i] = -1
	}
	return MyHashMap{hashmap}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hashmap[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.hashmap[key]
}

func (this *MyHashMap) Remove(key int) {
	this.hashmap[key] = -1
}

func main() {}
