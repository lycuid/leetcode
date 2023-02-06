// https://leetcode.com/problems/lfu-cache/description/
package main

type Block struct {
	key, value int
	back_ref   *List
	prev, next *Block
}

type List struct {
	id          int
	first, last *Block
	prev, next  *List
}

func (this *List) Attach(blk *Block) {
	if blk != nil {
		if blk.next, blk.back_ref = this.first, this; this.first == nil {
			this.first, this.last = blk, blk
		} else {
			this.first.prev = blk
		}
		this.first = blk
	}
}

func (this *List) Detach(blk *Block) *Block {
	if blk != nil {
		if blk.prev != nil {
			blk.prev.next = blk.next
		} else {
			this.first = blk.next
		}
		if blk.next != nil {
			blk.next.prev = blk.prev
		} else {
			this.last = blk.prev
		}
		blk.next, blk.prev, blk.back_ref = nil, nil, nil
	}
	return blk
}

type LFUCache struct {
	capacity   int
	blk_list   List
	blk_lookup map[int]*Block
}

func (this *LFUCache) Add(after, list *List) {
	if list != nil {
		if after.next != nil {
			after.next.prev = list
		}
		list.next, list.prev, after.next = after.next, after, list
	}
}

func (this *LFUCache) Remove(list *List) {
	if list != nil {
		if list.prev != nil {
			list.prev.next = list.next
		} else {
			this.blk_list.next = list
		}
		if list.next != nil {
			list.next.prev = list.prev
		}
		list.prev, list.next, list.first, list.last = nil, nil, nil, nil
	}
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		blk_list:   List{},
		blk_lookup: make(map[int]*Block),
	}
}

func (this *LFUCache) Advance(blk *Block) {
	list := blk.back_ref
	if list.Detach(blk); list.next == nil || list.next.id != list.id+1 {
		this.Add(list, &List{id: list.id + 1})
	}
	if list.next.Attach(blk); list.first == nil && list != &this.blk_list {
		this.Remove(list)
	}
}

func (this *LFUCache) Get(key int) int {
	if blk, found := this.blk_lookup[key]; found {
		this.Advance(blk)
		return blk.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if blk, found := this.blk_lookup[key]; !found {
		if this.capacity == 0 {
			if list := this.blk_list.next; list != nil {
				blk := list.Detach(list.last)
				this.capacity++
				if delete(this.blk_lookup, blk.key); list.first == nil {
					this.Remove(list)
				}
			}
		}
		if this.capacity > 0 {
			blk = &Block{key: key, value: value}
			this.blk_list.Attach(blk)
			this.blk_lookup[key], this.capacity = blk, this.capacity-1
			this.Advance(blk)
		}
	} else {
		blk.value = value
		this.Advance(blk)
	}
}

func main() {}
