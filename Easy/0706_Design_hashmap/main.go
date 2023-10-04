// https://leetcode.com/problems/design-hashmap/
package main

type Bucket struct {
	Next       *Bucket
	key, value int
}

type MyHashMap struct{ buckets [100]*Bucket }

func Constructor() MyHashMap { return MyHashMap{} }

func (this MyHashMap) get_bucket(key int) *Bucket {
	bucket := this.buckets[key%len(this.buckets)]
	for bucket != nil && bucket.key != key {
		bucket = bucket.Next
	}
	return bucket
}

func (this *MyHashMap) Put(key int, value int) {
	head := &this.buckets[key%len(this.buckets)]
	if bucket := this.get_bucket(key); bucket != nil {
		bucket.value = value
	} else {
		bucket = &Bucket{key: key, value: value, Next: *head}
		*head = bucket
	}
}

func (this *MyHashMap) Get(key int) int {
	if bucket := this.get_bucket(key); bucket != nil {
		return bucket.value
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	var previous, current *Bucket
	head := &this.buckets[key%len(this.buckets)]
	for current = *head; current != nil; previous, current = current, current.Next {
		if current.key == key {
			if previous == nil {
				*head = current.Next
			} else {
				previous.Next = current.Next
			}
			break
		}
	}
}

func main() {}
