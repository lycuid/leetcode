// https://leetcode.com/problems/single-threaded-cpu/
package main

type PriorityQueue struct {
	inner []int
	Less  func(i, j int) bool
}

func MakePQ(less func(i, j int) bool) PriorityQueue {
	return PriorityQueue{make([]int, 1), less}
}

func (this PriorityQueue) Empty() bool { return len(this.inner) == 1 }
func (this PriorityQueue) First() int  { return this.inner[1] }

func (this *PriorityQueue) Push(item int) {
	this.inner = append(this.inner, item)
	for i := len(this.inner) - 1; i > 1 && this.Less(this.inner[i], this.inner[i/2]); i /= 2 {
		this.inner[i], this.inner[i/2] = this.inner[i/2], this.inner[i]
	}
}

func (this *PriorityQueue) Pop() int {
	item, n := this.inner[1], len(this.inner)-1
	this.inner[1], this.inner = this.inner[n], this.inner[:n]
	for i, j := 1, 2; j < n; i, j = j, j*2 {
		if j+1 < n && this.Less(this.inner[j+1], this.inner[j]) {
			j++
		}
		if this.Less(this.inner[i], this.inner[j]) {
			break
		}
		this.inner[i], this.inner[j] = this.inner[j], this.inner[i]
	}
	return item
}

func getOrder(tasks [][]int) []int {
	order := make([]int, len(tasks))
	waiting := MakePQ(func(i, j int) bool { return tasks[i][0] < tasks[j][0] })
	ready := MakePQ(func(i, j int) bool {
		if tasks[i][1] == tasks[j][1] {
			return i < j
		}
		return tasks[i][1] < tasks[j][1]
	})
	for task := range tasks {
		waiting.Push(task)
	}
	for current_time, i := 0, 0; !waiting.Empty(); {
		current_time = tasks[waiting.First()][0]
		for !waiting.Empty() && tasks[waiting.First()][0] <= current_time {
			ready.Push(waiting.Pop())
		}
		for !ready.Empty() {
			t := ready.Pop()
			current_time, order[i], i = current_time+tasks[t][1], t, i+1
			for !waiting.Empty() && tasks[waiting.First()][0] <= current_time {
				ready.Push(waiting.Pop())
			}
		}
	}
	return order
}

func main() {}
